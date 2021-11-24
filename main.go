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

	"github.com/c-ollins/crabada/idlegame"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"

	IdleContractAddress = "0x82a85407BD612f52577909F4A58bfC6873f14DA8"

	TelegramChatID = 0

	TeamAvailable = "AVAILABLE"

	POLL_URL  = "https://idle-api.crabada.com/public/idle/mines?page=1&status=open&looter_address=0xed3428bcc71d3b0a43bb50a64ed774bec57100a8&can_loot=1&limit=8"
	LOOT_URL  = "https://idle-api.crabada.com/public/idle/mines?looter_address=%s&page=1&status=open&limit=8"
	TEAMS_URL = "https://idle-api.crabada.com/public/idle/teams?user_address=%s"

	GAS_API = "https://api.debank.com/chain/gas_price_dict_v2?chain=avax"
)

var (
	TelegramChat   = &tb.Chat{ID: TelegramChatID}
	MsgSendOptions = &tb.SendOptions{DisableWebPagePreview: true}

	settleRegex = regexp.MustCompile(`\/settle (.+)`)
	attackRegex = regexp.MustCompile(`\/attack (.+)`)

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

	client, err := ethclient.Dial("wss://api.avax.network/ext/bc/C/ws")
	if err != nil {
		log.Error(err)
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
			go et.raid()
			return
		case m.Text == "/loots":
			go et.sendActiveLoots(m)
			return
		case m.Text == "/settleall":
			go et.settleAll(false)
			return
		case m.Text == "/teams":
			go et.sendTeams(m)
			return
		}

		if matches := settleRegex.FindStringSubmatch(m.Text); len(matches) > 1 {
			gameID, err := strconv.Atoi(matches[1])
			if err != nil {
				b.Reply(m, err.Error())
				return
			}

			go et.settleGame(int64(gameID))
		} else if matches := attackRegex.FindStringSubmatch(m.Text); len(matches) > 1 {
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
	go et.queAttacks()
	if et.isAuto {
		go et.auto()
	}
	go et.gasUpdate()
	b.Start()
}

func (et *etubot) raid() {
	teams, err := et.allTeams()
	if err != nil {
		et.sendError(fmt.Errorf("error finding teams team:%v", err))
		return
	}

	queue := 0
	for _, team := range teams {
		if team.Status == TeamAvailable {
			queue++
			et.attackCh <- team
		}
	}
	if queue > 0 {
		et.bot.Send(TelegramChat, fmt.Sprintf("Queued %d teams for attack.", queue))
	} else if !et.isAuto {
		et.bot.Send(TelegramChat, "All teams are busy.")
	}
}

func (et *etubot) queAttacks() {
	log.Info("Attacks queue active")
	for team := range et.attackCh {

		if !et.teamIsAvailable(team.ID) {
			lg := fmt.Sprintf("Cannot attack, team %d is not available", team.ID)
			log.Info(lg)
			et.bot.Send(TelegramChat, lg)
			continue
		}

		et.pollGamesAndAttack(team)
	}
}

func (et *etubot) pollGamesAndAttack(team *Team) {
	errorCount := 0

	for {
		var response GamesResponse
		err := makeRequest(POLL_URL, &response)
		if err != nil {
			log.Error("error polling games:", err)
			return
		}

		if response.ErrorCode != "" {
			log.Error("Error polling games code: %s, message %s", response.ErrorCode, response.Message)
			if response.ErrorCode == INTERNAL_SERVER_ERROR {
				continue
			}

			return
		}

		log.Infof("Found %d games", response.TotalRecord)
		var target *Game
		for _, game := range response.GamesResult.Games {

			if team.Strength > game.DefensePoint && time.Since(time.Unix(game.StartTime, 0)) < (10*time.Second) {
				target = &game
				break
			}
		}

		if target == nil {
			time.Sleep(200)
			continue
		}

		log.Infof("Game: %d, opponent strength: %d, team strength: %d, start time:%s", target.ID, target.DefensePoint, team.Strength, time.Since(time.Unix(target.StartTime, 0)).Truncate(time.Second))

		err = et.attack(target, team)
		if err != nil {
			errorCount++
			if errorCount >= 10 {
				et.bot.Send(TelegramChat, fmt.Sprintf("More than 10 errors while trying to attack using team %d. %v", team.ID, err))
				return
			}

			log.Error("error attacking:", err)
			continue
		}

		break
	}
}

func (et *etubot) attack(game *Game, team *Team) error {

	lg := fmt.Sprintf("Attacking %d using %d, strength advantage: %d.", game.ID, team.ID, team.Strength-game.DefensePoint)
	log.Info(lg)
	auth, err := et.txAuth(team.Wallet)
	if err != nil {
		return err
	}

	tx, err := et.idleContract.Attack(auth, big.NewInt(game.ID), big.NewInt(team.ID))
	if err != nil {
		return err
	}

	log.Info("Attack tx hash:", tx.Hash())
	// wait for receipt
	waitStart := time.Now()
	for {
		receipt, err := et.client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			if err != ethereum.NotFound {
				log.Error("error:", err)
			}
			time.Sleep(5 * time.Second)
			if time.Since(waitStart) > 2*time.Minute {
				return fmt.Errorf("transaction failed on: %d, did not get receipt after 2 minutes", game.ID)
			}
			continue
		}

		log.Info(receipt)
		if receipt.Status == types.ReceiptStatusSuccessful {
			et.bot.Send(TelegramChat, fmt.Sprintf("Game #%d attack successful by team #%d.\nhttps://snowtrace.io/tx/%s", game.ID, team.ID, tx.Hash().String()), MsgSendOptions)
			return nil
		}

		log.Info("Attack failed, retrying")
		return fmt.Errorf("transaction failed on: %d", game.ID)
	}
}

func (et *etubot) sendError(err error) {
	log.Error(err)
	et.bot.Send(TelegramChat, err)
}

func (et *etubot) settleAll(isAuto bool) {
	games, err := et.activeLoots()
	if err != nil {
		et.sendError(fmt.Errorf("error fetching active loots: %v", err))
		return
	}

	totalSettled := 0

	for _, game := range games {
		tm := time.Unix(game.StartTime, 0)
		settle := tm.Add(time.Duration(1) * time.Hour).Add(time.Duration(5) * time.Minute)
		if time.Now().After(settle) {
			totalSettled++
			et.settleGame(game.ID)
		}
	}

	if totalSettled == 0 && !isAuto {
		et.bot.Send(TelegramChat, "No games ready to be settled.")
	}
}

func (et *etubot) settleGame(gameID int64) {
	log.Info("Settling game", gameID)
	et.bot.Send(TelegramChat, fmt.Sprintf("Settling game #%d", gameID))
	team, err := et.findMyLootTeam(gameID)
	if err != nil {
		et.sendError(fmt.Errorf("error finding loot team:%v", err))
		return
	}

	auth, err := et.txAuth(team.Wallet)
	if err != nil {
		et.sendError(fmt.Errorf("error creating tx auth: %v", err))
		return
	}

	tx, err := et.idleContract.SettleGame(auth, big.NewInt(gameID))
	if err != nil {
		et.sendError(fmt.Errorf("error sending settle tx: %v", err))
		return
	}

	log.Info("Settle hash:", tx.Hash())
	// wait for receipt
	waitStart := time.Now()
	for {
		log.Info("Checking for receipt")
		_, err := et.client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			if err != ethereum.NotFound {
				log.Error("error:", err)
			}
			if time.Since(waitStart) > 2*time.Minute {
				et.bot.Send(TelegramChat, fmt.Sprintf("Game #%d Team #%d has been settled, but didnt not get confirmed in 1 minute.\nhttps://snowtrace.io/tx/%s", gameID, team.ID, tx.Hash().String()), MsgSendOptions)
				return
			}
			time.Sleep(5 * time.Second)
			continue
		}
		et.bot.Send(TelegramChat, fmt.Sprintf("Game #%d Team #%d has been settled.\nhttps://snowtrace.io/tx/%s", gameID, team.ID, tx.Hash().String()), MsgSendOptions)
		break
	}
}

func (et *etubot) findMyLootTeam(gameID int64) (*Team, error) {
	var response GameResponse
	err := makeRequest(fmt.Sprintf("https://idle-api.crabada.com/public/idle/mine/%d", gameID), &response)
	if err != nil {
		return nil, fmt.Errorf("error fetching game by id: %v", err)
	}

	if response.ErrorCode != "" {
		return nil, fmt.Errorf("error fetching game by id: %s, message: %s", response.ErrorCode, response.Message)
	}

	// TODO: verify team

	return &Team{ID: response.AttackTeamID, Strength: response.AttackPoint, Wallet: response.AttackTeamOwner}, nil
}

func (et *etubot) txAuth(address string) (*bind.TransactOpts, error) {
	nonce, err := et.client.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	// gasPrice, err := et.client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return nil, err
	// }

	auth, err := bind.NewKeyedTransactorWithChainID(et.privateKey[strings.ToLower(address)], big.NewInt(43114))
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(200000) // in units // TODO
	// auth.GasPrice = big.NewInt(0).Sub(gasPrice, big.NewInt(50000000000)) // add 30 gwei
	et.gasMu.RLock()
	auth.GasPrice = big.NewInt(0).Sub(et.gasPrice, big.NewInt(20000000000))
	et.gasMu.RUnlock()
	log.Info("Using gas:", big.NewInt(0).Div(auth.GasPrice, big.NewInt(1e9)))

	return auth, nil
}

func (et *etubot) activeLoots() ([]Game, error) {
	var games []Game
	for i := 0; i < len(wallets); i++ {

		var response GamesResponse
		err := makeRequest(fmt.Sprintf(LOOT_URL, wallets[i]), &response)
		if err != nil {
			log.Error("Error fetching active loots:", err)
			return nil, err
		}

		if response.ErrorCode != "" {
			return nil, fmt.Errorf("error fetching game by id: %s, message: %s", response.ErrorCode, response.Message)
		}

		games = append(games, response.Games...)
	}

	return games, nil
}

func (et *etubot) sendActiveLoots(msg *tb.Message) {

	games, err := et.activeLoots()
	if err != nil {
		et.sendError(fmt.Errorf("error fetching active loots: %v", err))
		return
	}

	sb := fmt.Sprintf("%d active loots 💰🤑💰\n-------------------------\n", len(games))
	for _, game := range games {
		tm := time.Unix(game.StartTime, 0)

		settle := tm.Add(time.Duration(1) * time.Hour).Add(time.Duration(5) * time.Minute)
		lootSummary := "💰 Loot\n"
		lootSummary += fmt.Sprintf("Game: %d\n", game.ID)
		lootSummary += fmt.Sprintf("Team: %d\n", game.AttackTeamID)
		lootSummary += fmt.Sprintf("Account: %s\n", game.AttackTeamOwner[:7])
		if time.Now().After(settle) {
			lootSummary += "Ready: yes\n"
		} else {
			lootSummary += fmt.Sprintf("Settle in: %s\n", settle.Sub(time.Now()))
		}

		sb += "\n" + lootSummary
	}

	et.bot.Reply(msg, sb)
}

func (et *etubot) allTeams() ([]*Team, error) {
	var teams []*Team
	for i := 0; i < len(wallets); i++ {
		var response TeamsResponse
		err := makeRequest(fmt.Sprintf(TEAMS_URL, wallets[i]), &response)
		if err != nil {
			log.Error("Error fetching teams:", err)
			return nil, err
		}

		if response.ErrorCode != "" {
			return nil, fmt.Errorf("error fetching game by id: %s, message: %s", response.ErrorCode, response.Message)
		}

		if response.TotalRecord > 0 {
			teams = append(teams, response.Teams...)
		}
	}

	return teams, nil
}

func (et *etubot) teamForID(teamID int64) (*Team, error) {
	teams, err := et.allTeams()
	if err != nil {
		return nil, err
	}

	for _, team := range teams {
		if team.ID == teamID {
			return team, nil
		}
	}

	return nil, fmt.Errorf("team does not exist in wallets")
}

func (et *etubot) teamIsAvailable(teamID int64) bool {
	team, err := et.teamForID(teamID)
	if err != nil {
		et.sendError(fmt.Errorf("error fetching teams: %v", err))
		return false
	}

	if team.Status == "AVAILABLE" {
		return true
	}

	return false
}

func (et *etubot) sendTeams(msg *tb.Message) {

	teams, err := et.allTeams()
	if err != nil {
		et.sendError(fmt.Errorf("error fetching teams: %v", err))
		return
	}

	sb := fmt.Sprintf("%d teams\n---------------------\n", len(teams))

	for _, team := range teams {
		teamSummary := fmt.Sprintf("ID: %d\n", team.ID)
		teamSummary += fmt.Sprintf("Strength: %d\n", team.Strength)
		teamSummary += fmt.Sprintf("Account: %s\n", team.Wallet[:7])
		teamSummary += fmt.Sprintf("Status: %s\n", strings.ToLower(team.Status))

		sb += "\n" + teamSummary
	}

	et.bot.Reply(msg, sb)
}

func (et *etubot) gas(msg *tb.Message) {

	et.gasMu.RLock()
	gasPrice := et.gasPrice
	et.gasMu.RUnlock()

	et.bot.Reply(msg, fmt.Sprintf("%d gwei", big.NewInt(0).Div(gasPrice, big.NewInt(1e9))))
}

func (et *etubot) gasUpdate() {
	fmt.Println("Gas update running")
	for {
		err := et.updateGasPrice()
		if err != nil {
			et.sendError(fmt.Errorf("err updating gas: %v", err))
		}
		time.Sleep(10 * time.Second)
	}
}

func (et *etubot) updateGasPrice() error {
	var response GasResponse
	err := makeRequest(GAS_API, &response)
	if err != nil {
		return fmt.Errorf("error fetching gas: %v", err)
	}

	gasPrice := big.NewInt(response.Data.Fast.Price)
	et.gasMu.Lock()
	et.gasPrice = gasPrice
	et.gasMu.Unlock()

	return nil
}

func (et *etubot) auto() {
	et.bot.Send(TelegramChat, "Etubot running on auto.")
	for {
		log.Info("Auto running")
		// settle ready games
		et.settleAll(true)
		et.raid()

		time.Sleep(5 * time.Minute)
	}
}
