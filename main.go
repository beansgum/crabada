package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/c-ollins/crabada/idlegame"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	InternalServerError = "INTERNAL_SERVER_ERROR"

	IdleContractAddress = "0x82a85407BD612f52577909F4A58bfC6873f14DA8"

	TelegramChatID = 0

	TeamAvailable = "AVAILABLE"

	LootURL  = "https://idle-api.crabada.com/public/idle/mines?looter_address=%s&page=1&status=open&limit=8"
	TeamsURL = "https://idle-api.crabada.com/public/idle/teams?user_address=%s"
	CrabsURL = "https://idle-api.crabada.com/public/idle/crabadas/can-join-team?user_address=%s"

	GasAPI = "https://api.debank.com/chain/gas_price_dict_v2?chain=avax"
)

var (
	TelegramChat   = &tb.Chat{ID: TelegramChatID}
	MsgSendOptions = &tb.SendOptions{DisableWebPagePreview: true}

	settleRegex = regexp.MustCompile(`\/settle (.+)`)
	attackRegex = regexp.MustCompile(`\/attack (.+)`)

	actionReinforceDefense = "reinforce-defense"
	actionReinforceAttack  = "reinforce-attack"
	actionAttack           = "attack"

	processIntervals = 30 * time.Minute

	wallets = []string{"0xed3428BcC71d3B0a43Bb50a64ed774bEc57100a8", "0xf91fF01b9EF0d83D0bBd89953d53504f099A3DFf"}
)

func main() {
	et := etubot{
		isAuto:     true,
		attackCh:   make(chan *Team, 5),
		privateKey: make(map[string]*ecdsa.PrivateKey),
	}
	et.start()
}

type etubot struct {
	bot    *tb.Bot
	isAuto bool

	gasPrice *big.Int
	gasMu    sync.RWMutex

	attackCh chan *Team

	client       *ethclient.Client
	idleContract *idlegame.Idlegame
	privateKey   map[string]*ecdsa.PrivateKey
}

func (et *etubot) start() {

	initLogRotator("logs.txt")

	privateKey, err := crypto.HexToECDSA(os.Getenv("BOT_PRIVATE"))
	if err != nil {
		log.Error(err)
		return
	}
	et.privateKey[strings.ToLower("0xed3428BcC71d3B0a43Bb50a64ed774bEc57100a8")] = privateKey

	privateKey2, err := crypto.HexToECDSA(os.Getenv("BOT_PRIVATE2"))
	if err != nil {
		log.Error(err)
		return
	}
	et.privateKey[strings.ToLower("0xf91fF01b9EF0d83D0bBd89953d53504f099A3DFf")] = privateKey2

	log.Info("Connecting to infura")
	client, err := ethclient.Dial("wss://speedy-nodes-nyc.moralis.io//avalanche/mainnet/ws")
	if err != nil {
		log.Error("ethclient:", err)
		return
	}
	et.client = client

	log.Info("Connected to infura")

	address := common.HexToAddress(IdleContractAddress)
	idleContract, err := idlegame.NewIdlegame(address, client)
	if err != nil {
		log.Error(err)
		return
	}

	et.idleContract = idleContract

	b, err := tb.NewBot(tb.Settings{
		Token:  "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Error(err)
		return
	}
	et.bot = b

	b.Handle(tb.OnText, func(m *tb.Message) {
		switch {
		case m.Text == "/ping":
			b.Reply(m, "pong!")
			return
		case m.Text == "/gas":
			go et.gas(m)
			return
		case m.Text == "/raid":
			if et.isAuto {
				et.bot.Send(TelegramChat, "cmd disabled in auto.")
				return
			}
			go et.raid()
			return
		case m.Text == "/loots":
			go et.sendActiveLoots(m)
			return
		case m.Text == "/settleall":
			if et.isAuto {
				et.bot.Send(TelegramChat, "cmd disabled in auto.")
				return
			}
			go et.settleAll(false)
			return
		case m.Text == "/teams":
			go et.sendTeams(m)
			return
		}

		if matches := settleRegex.FindStringSubmatch(m.Text); len(matches) > 1 {
			if et.isAuto {
				et.bot.Send(TelegramChat, "cmd disabled in auto.")
				return
			}

			gameID, err := strconv.Atoi(matches[1])
			if err != nil {
				b.Reply(m, err.Error())
				return
			}

			go et.settleGame(int64(gameID))
		} else if matches := attackRegex.FindStringSubmatch(m.Text); len(matches) > 1 {
			if et.isAuto {
				et.bot.Send(TelegramChat, "cmd disabled in auto.")
				return
			}

			teamID, err := strconv.Atoi(matches[1])
			if err != nil {
				b.Reply(m, err.Error())
				return
			}

			go func() {
				team, err := et.teamForID(int64(teamID))
				if err != nil {
					if err != nil {
						b.Reply(m, fmt.Sprintf("Could not find team:%v", err))
						return
					}
				}

				et.attackCh <- team
				b.Reply(m, fmt.Sprintf("Attacking using team #%d", team.ID))
			}()
		}
	})

	log.Info("Bot running")
	err = et.updateGasPrice()
	if err != nil {
		log.Error(err)
		return
	}

	// go et.watchStartGame()
	// go et.queAttacks()
	if et.isAuto {
		go et.auto()
	}
	go et.gasUpdate()
	b.Start()

	select {}
}

func (et *etubot) sendError(err error) {
	log.Error(err)
	// et.bot.Send(TelegramChat, err.Error())
}

func (et *etubot) txAuth(address string, addGas bool) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(et.privateKey[strings.ToLower(address)], big.NewInt(43114))
	if err != nil {
		return nil, fmt.Errorf("error creating transactor: %v", err)
	}

	limit := big.NewInt(210000000000) //210gwei

	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(200000) // in units
	et.gasMu.RLock()
	if addGas {
		auth.GasPrice = big.NewInt(0).Add(et.gasPrice, big.NewInt(100000000000)) //add 100 gwei
	} else {
		auth.GasPrice = et.gasPrice
	}
	et.gasMu.RUnlock()

	if auth.GasPrice.Cmp(limit) >= 0 {
		return nil, fmt.Errorf("cannot make tx, gas too high")
	}

	return auth, nil
}

func (et *etubot) updateGasPrice() error {
	var response GasResponse
	err := makeRequest(GasAPI, &response)
	if err != nil {
		return fmt.Errorf("error fetching gas: %v", err)
	}

	// fmt.Printf("New gas price: %d\n", int64(response.Data.Fast.Price))

	gasPrice := big.NewInt(int64(response.Data.Fast.Price))
	et.gasMu.Lock()
	et.gasPrice = gasPrice
	et.gasMu.Unlock()

	return nil
}
