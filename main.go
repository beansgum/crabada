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
	IdleContractAddress = "0x82a85407BD612f52577909F4A58bfC6873f14DA8"

	TelegramChatID = 0

	TeamAvailable = "AVAILABLE"

	GasAPI = "https://api.debank.com/chain/gas_price_dict_v2?chain=avax"
)

var (
	TelegramChat   = &tb.Chat{ID: TelegramChatID}
	MsgSendOptions = &tb.SendOptions{DisableWebPagePreview: true}

	settleRegex = regexp.MustCompile(`\/settle (.+)`)
	attackRegex = regexp.MustCompile(`\/attack (.+)`)
	watchRegex  = regexp.MustCompile(`\/watch (.+)`)

	processIntervals = 30 * time.Minute

	reinforcementCrabs = make(map[string][]int64)
)

func main() {
	reinforcementCrabs["0xed3428bcc71d3b0a43bb50a64ed774bec57100a8"] = []int64{8872, 8869, 8876, 8873, 8976, 8877}
	reinforcementCrabs["0xf91ff01b9ef0d83d0bbd89953d53504f099a3dff"] = []int64{8874, 8870, 8871, 8875, 8363, 8881}
	reinforcementCrabs["0x303de8234c60c146902f3e6f340722e41595667b"] = []int64{9390, 9637, 9393, 9387, 2584}
	et := etubot{
		isAuto:     true,
		attackCh:   make(chan *Team, 5),
		privateKey: make(map[string]*ecdsa.PrivateKey),
		games:      make(map[int64]int),
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

	games   map[int64]int
	gamesMu sync.Mutex
}

func (et *etubot) start() {

	initLogRotator("logs.txt")

	privateKey, err := crypto.HexToECDSA(os.Getenv("BOT_PRIVATE"))
	if err != nil {
		log.Error(err)
		return
	}
	et.privateKey["0xed3428bcc71d3b0a43bb50a64ed774bec57100a8"] = privateKey

	privateKey2, err := crypto.HexToECDSA(os.Getenv("BOT_PRIVATE2"))
	if err != nil {
		log.Error(err)
		return
	}
	et.privateKey["0xf91ff01b9ef0d83d0bbd89953d53504f099a3dff"] = privateKey2

	privateKey3, err := crypto.HexToECDSA(os.Getenv("BOT_PRIVATE3"))
	if err != nil {
		log.Error(err)
		return
	}
	et.privateKey["0x303de8234c60c146902f3e6f340722e41595667b"] = privateKey3

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
		case m.Text == "/reinforce":
			go et.reinforceAttacks()
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
		} else if matches := watchRegex.FindStringSubmatch(m.Text); len(matches) > 1 {
			gameID, err := strconv.Atoi(matches[1])
			if err != nil {
				b.Reply(m, err.Error())
				return
			}

			et.gamesMu.Lock()
			et.games[int64(gameID)] = 1
			et.gamesMu.Unlock()
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
	err = et.addActiveGames()
	if err != nil {
		log.Error(err)
		return
	}
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
