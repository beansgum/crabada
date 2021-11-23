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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	IdleContractAddress = "0x82a85407BD612f52577909F4A58bfC6873f14DA8"

	TelegramChatID = 0

	POLL_URL  = "https://idle-api.crabada.com/public/idle/mines?page=1&status=open&looter_address=0xed3428bcc71d3b0a43bb50a64ed774bec57100a8&can_loot=1&limit=8"
	LOOT_URL  = "https://idle-api.crabada.com/public/idle/mines?looter_address=0xed3428bcc71d3b0a43bb50a64ed774bec57100a8&page=1&status=open&limit=8"
	TEAMS_URL = "https://idle-api.crabada.com/public/idle/teams?user_address=%s"
)

var (
	TelegramChat   = tb.Chat{ID: TelegramChatID}
	MsgSendOptions = &tb.SendOptions{DisableWebPagePreview: true}

	settleRegex = regexp.MustCompile(`\/settle (.+)`)

	wallets = []string{"0xed3428BcC71d3B0a43Bb50a64ed774bEc57100a8", "0xf91fF01b9EF0d83D0bBd89953d53504f099A3DFf"}
)

func main() {
	et := etubot{}
	et.start()
}

type etubot struct {
	mu sync.RWMutex

	bot *tb.Bot

	client       *ethclient.Client
	idleContract *idlegame.Idlegame
	privateKey   *ecdsa.PrivateKey
}

func (et *etubot) start() {

	initLogRotator("logs.txt")

	privateKey, err := crypto.HexToECDSA(os.Getenv("BOT_PRIVATE"))
	if err != nil {
		log.Error(err)
		return
	}
	et.privateKey = privateKey

	client, err := ethclient.Dial("wss://api.avax.network/ext/bc/C/ws")
	if err != nil {
		log.Error(err)
		return
	}
	et.client = client

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Error(err)
	// 	return
	// }

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
		case m.Text == "/loots":
			go et.sendActiveLoots()
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

			go et.settleGame(m, int64(gameID))
		}

	})

	fmt.Println()

	log.Info("Bot running")
	// et.pollGamesAndAttack(nil)
	b.Start()
}

func (et *etubot) pollGamesAndAttack(team *Team) {
	// fmt.Printf("Polling active games for  team #%d strength: %d", team.ID, team.Strength)
	var response GamesResponse
	err := makeRequest(POLL_URL, &response)
	if err != nil {
		fmt.Println("Error polling games:", err)
		return
	}

	if response.ErrorCode != "" {
		fmt.Println("Error polling games code:", response.ErrorCode, "message:"+response.Message)
		return
	}

	fmt.Printf("Found %d games\n", response.TotalRecord)
	if response.TotalRecord > 0 {
		for _, game := range response.GamesResult.Games {
			fmt.Println("Game:", game.ID, "opponent strength:", game.DefensePoint, "team strength:", team.Strength)
			if team.Strength > game.DefensePoint {

			}
		}
	}
}

func (et *etubot) attack() {

}

func (et *etubot) sendError(err error) {
	log.Error(err)
	et.bot.Send(&TelegramChat, err)
}

func (et *etubot) settleGame(msg *tb.Message, gameID int64) {
	fmt.Println("Settling game", gameID)
	et.bot.Reply(msg, fmt.Sprintf("Settling game #%d", gameID))
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

	fmt.Println("Settle hash:", tx.Hash())
	// wait for receipt
	go func() {
		for {
			fmt.Println("Checking for receipt")
			_, err := et.client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				if err != ethereum.NotFound {
					fmt.Println("error:", err)
				}
				time.Sleep(5 * time.Second)
				continue
			}
			et.bot.Reply(msg, fmt.Sprintf("Game #%d Team #%d has been settled.\nhttps://snowtrace.io/tx/%s", gameID, team.ID, tx.Hash().String()), MsgSendOptions)
			break
		}
	}()
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

	gasPrice, err := et.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(et.privateKey, big.NewInt(43114))
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(145000) // in units // TODO
	auth.GasPrice = gasPrice

	return auth, nil
}

func (et *etubot) sendActiveLoots() {
	var response GamesResponse
	err := makeRequest(LOOT_URL, &response)
	if err != nil {
		fmt.Println("Error fetching active loots:", err)
		return
	}

	if response.ErrorCode != "" {
		fmt.Println("Error polling games code:", response.ErrorCode, "message:"+response.Message)
		return
	}

	if response.TotalRecord > 0 {
		sb := fmt.Sprintf("%d active loots 💰🤑💰\n-------------------------\n", response.TotalRecord)

		et.bot.Send(&TelegramChat, sb)
		for _, game := range response.GamesResult.Games {
			tm := time.Unix(game.StartTime, 0)

			settle := tm.Add(time.Duration(1) * time.Hour).Add(time.Duration(5) * time.Minute)
			lootSummary := "💰 Loot\n"
			lootSummary += fmt.Sprintf("Game: %d\n", game.ID)
			lootSummary += fmt.Sprintf("Team: %d\n", game.AttackTeamID)
			if time.Now().After(settle) {
				lootSummary += "Ready: yes"
			} else {
				lootSummary += fmt.Sprintf("Settle in: %s\n", time.Until(settle))
			}

			et.bot.Send(&TelegramChat, lootSummary)
		}
	}
}

func (et *etubot) allTeams() ([]Team, error) {
	var teams []Team
	for i := 0; i < len(wallets); i++ {
		var response TeamsResponse
		err := makeRequest(fmt.Sprintf(TEAMS_URL, wallets[i]), &response)
		if err != nil {
			fmt.Println("Error fetching teams:", err)
			return nil, err
		}

		if response.ErrorCode != "" {
			fmt.Println("Error fetching teams code:", response.ErrorCode, "message:"+response.Message)
			return nil, err
		}

		if response.TotalRecord > 0 {
			teams = append(teams, response.Teams...)
		}
	}

	return teams, nil
}

func (et *etubot) teamIsAvailable(teamID int) bool {
	teams, err := et.allTeams()
	if err != nil {
		et.sendError(fmt.Errorf("error fetching teams: %v", err))
		return false
	}

	for _, team := range teams {
		if team.ID == teamID {
			if team.Status == "AVAILABLE" {
				return true
			}
		}
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
