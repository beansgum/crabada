package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/c-ollins/crabada/crabcaller"
	"github.com/c-ollins/crabada/idlegame"
	"github.com/c-ollins/crabada/traderjoe"
	"github.com/c-ollins/crabada/tus"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	RaidGasLimit = big.NewInt(600000000000)
	RaidGasMin   = big.NewInt(300000000000)
	RaidGasExtra = big.NewInt(70000000000)
	GasLimit     = big.NewInt(310000000000)

	MinBalance = ToWei(0.4)
)

func main() {
	reinforcementCrabs["0xed3428bcc71d3b0a43bb50a64ed774bec57100a8"] = []int64{8872, 8869, 8876, 8873, 8976, 8877}
	reinforcementCrabs["0xf91ff01b9ef0d83d0bbd89953d53504f099a3dff"] = []int64{8874, 8870, 8871, 8875, 8363, 8881}
	reinforcementCrabs["0x303de8234c60c146902f3e6f340722e41595667b"] = []int64{9390, 9637, 9393, 9387, 2584, 11448}
	et := etubot{
		isAuto:        true,
		lootingActive: true,
		raidGasPrice:  RaidGasMin,
		attackCh:      make(chan *Team, 5),
		privateKey:    make(map[string]*ecdsa.PrivateKey),
		games:         make(map[int64]int),
	}
	et.start()
}

type etubot struct {
	bot           *tb.Bot
	isAuto        bool
	lootingActive bool

	gasPrice *big.Int
	gasMu    sync.RWMutex

	raidGasPrice   *big.Int
	raidGasMu      sync.RWMutex
	raidLastUpdate time.Time

	attackCh chan *Team

	client       *ethclient.Client
	idleContract *idlegame.Idlegame
	crabCaller   *crabcaller.Crabcaller
	traderJoe    *traderjoe.Traderjoe
	tus          *tus.Tus
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

	err = et.connect()
	if err != nil {
		log.Errorf("error connecting: %v", err)
		return
	}

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
		case m.Text == "/reconnect":
			go et.reconnect()
			return
		case m.Text == "/gas":
			go et.gas(m)
			return
		case m.Text == "/raidgas":
			go et.raidGas(m)
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
		case m.Text == "/pauseattacks":
			et.lootingActive = false
			et.bot.Send(TelegramChat, "looting paused")
			return
		case m.Text == "/startattacks":
			et.lootingActive = true
			et.bot.Send(TelegramChat, "looting active")
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
		log.Error("error adding active games:", err)
		return
	}
	if et.isAuto {
		go et.auto()
	}
	go et.gasUpdate()
	go et.watchRaidGas()
	b.Start()

	select {}
}

func (et *etubot) reconnect() {
	et.client.Close()
	err := et.connect()
	if err != nil {
		et.bot.Send(TelegramChat, fmt.Sprintf("error reconnecting: %v", err))
	}
}

func (et *etubot) connect() error {
	log.Info("Connecting to infura")
	//ws://127.0.0.1:9650/ext/bc/C/ws
	//wss://speedy-nodes-nyc.moralis.io//avalanche/mainnet/ws
	client, err := ethclient.Dial("ws://127.0.0.1:9650/ext/bc/C/ws")
	if err != nil {
		return err
	}

	log.Info("connected client")

	et.client = client

	log.Info("Connected to infura")

	crabadaAddress := common.HexToAddress(IdleContractAddress)
	idleContract, err := idlegame.NewIdlegame(crabadaAddress, client)
	if err != nil {
		return err
	}

	crabCallerAddress := common.HexToAddress("0xCe2139165a38BFD2D867B16D6cc1aB2a6171191a")
	crabCaller, err := crabcaller.NewCrabcaller(crabCallerAddress, client)
	if err != nil {
		return err
	}

	traderJoeAddress := common.HexToAddress("0x60ae616a2155ee3d9a68541ba4544862310933d4")
	traderJoe, err := traderjoe.NewTraderjoe(traderJoeAddress, client)
	if err != nil {
		return err
	}

	tusAddress := common.HexToAddress("0xf693248F96Fe03422FEa95aC0aFbBBc4a8FdD172")
	tusToken, err := tus.NewTus(tusAddress, client)
	if err != nil {
		return err
	}

	et.crabCaller = crabCaller
	et.idleContract = idleContract
	et.traderJoe = traderJoe
	et.tus = tusToken
	return nil
}

func (et *etubot) sendError(err error) {
	log.Error(err)
	// et.bot.Send(TelegramChat, err.Error())
}

func (et *etubot) txAuth(address string, raidGas bool) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(et.privateKey[strings.ToLower(address)], big.NewInt(43114))
	if err != nil {
		return nil, fmt.Errorf("error creating transactor: %v", err)
	}

	nonce, err := et.client.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return nil, fmt.Errorf("error creating transactor: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	et.gasMu.RLock()

	if et.gasPrice.Cmp(GasLimit) >= 0 {
		et.gasMu.RUnlock()
		return nil, fmt.Errorf("cannot make tx, gas too high")
	}
	auth.GasPrice = et.gasPrice
	et.gasMu.RUnlock()

	if raidGas {

		et.raidGasMu.RLock()
		auth.GasTipCap = big.NewInt(0).Add(et.raidGasPrice, RaidGasExtra)
		auth.GasFeeCap = auth.GasTipCap
		auth.GasPrice = nil
		et.raidGasMu.RUnlock()

		if auth.GasTipCap.Cmp(RaidGasLimit) >= 0 {
			return nil, fmt.Errorf("cannot make tx, raid gas too high: %v gwei", ToGwei(auth.GasTipCap))
		}
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

func (et *etubot) shouldSellTus(address string) (bool, error) {
	balance, err := et.client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return false, err
	}

	if balance.Cmp(MinBalance) == 1 {
		return false, nil
	}

	callOpts := &bind.CallOpts{Context: context.Background()}
	tusBalance, err := et.tus.BalanceOf(callOpts, common.HexToAddress(address))
	if err != nil {
		return false, err
	}

	if tusBalance.Cmp(ToWei(int64(600))) == -1 {
		return false, fmt.Errorf("cannot sell tus for %s, not enough tus balance", address)
	}

	return true, nil
}

func (et *etubot) sellTus(address string) error {
	et.bot.Send(TelegramChat, fmt.Sprintf("Selling 500 $TUS for AVAX using account %s", address), MsgSendOptions)

	opts, err := et.txAuth(address, false)
	if err != nil {
		return err
	}

	routerPath := []common.Address{
		common.HexToAddress("0xf693248F96Fe03422FEa95aC0aFbBBc4a8FdD172"),
		common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7"),
	}
	deadline := time.Now().Add(4*time.Minute).UnixNano() / int64(time.Millisecond)
	tx, err := et.traderJoe.SwapExactTokensForAVAX(opts, ToWei(int64(500)), ToWei(float64(0.5)), routerPath, common.HexToAddress(address), big.NewInt(deadline))
	if err != nil {
		return err
	}

	waitStart := time.Now()
	for {
		receipt, err := et.client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			if err != ethereum.NotFound {
				log.Error("error:", err)
			}
			time.Sleep(5 * time.Second)
			if time.Since(waitStart) > 2*time.Minute {
				return fmt.Errorf("transaction failed when selling tus, did not get receipt after 2 minutes")
			}
			continue
		}

		if receipt.Status == types.ReceiptStatusSuccessful {
			txt := fmt.Sprintf("500 $TUS sold for AVAX.\nhttps://snowtrace.io/tx/%s", tx.Hash())
			et.bot.Send(TelegramChat, txt, MsgSendOptions)
		} else {
			log.Info("sell tus failed, retrying")
		}

		break
	}
	return nil
}
