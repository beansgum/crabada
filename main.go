package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	Network        = "mainnet" // rinkeby | mainnet
	TelegramChatID = 0

	POLL_URL = "https://idle-api.crabada.com/public/idle/mines?page=1&status=open&looter_address=0xed3428bcc71d3b0a43bb50a64ed774bec57100a8&can_loot=1&limit=8"
	LOOT_URL = "https://idle-api.crabada.com/public/idle/mines?looter_address=0xed3428bcc71d3b0a43bb50a64ed774bec57100a8&page=1&status=open&limit=8"
)

var TelegramChat = tb.Chat{ID: TelegramChatID}

func main() {
	et := etubot{}
	et.start()
}

type etubot struct {
	mu sync.RWMutex
	// orders map[string]*LimitOrder // [Contract Address]*LimitOrder

	bot *tb.Bot

	client *ethclient.Client
}

func (et *etubot) start() {

	initLogRotator("logs.txt")
	client, err := ethclient.Dial("wss://" + Network + ".infura.io/ws/v3/")
	if err != nil {
		log.Error(err)
		return
	}
	et.client = client

	log.Info("Connected to infura")

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
		switch m.Text {
		case "/ping":
			b.Reply(m, "pong!")
			return
		case "/loots":
			et.sendActiveLoots()
			return

		}
	})

	log.Info("Bot running")
	// et.pollGamesAndAttack(nil)
	b.Start()
}

func (et *etubot) pollGamesAndAttack(team *Team) {
	// fmt.Printf("Polling active games for  team #%d strength: %d", team.ID, team.Strength)
	var response GameResponse
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
		for _, game := range response.GameResult.Games {
			fmt.Println("Game:", game.ID, "opponent strength:", game.DefensePoint, "team strength:", team.Strength)
			if team.Strength > game.DefensePoint {

			}
		}
	}
}

func (et *etubot) attack() {

}

func (et *etubot) sendActiveLoots() {
	var response GameResponse
	err := makeRequest(LOOT_URL, &response)
	if err != nil {
		fmt.Println("Error polling games:", err)
		return
	}

	if response.ErrorCode != "" {
		fmt.Println("Error polling games code:", response.ErrorCode, "message:"+response.Message)
		return
	}

	if response.TotalRecord > 0 {
		sb := fmt.Sprintf("%d active loots 💰🤑💰\n-------------------------\n", response.TotalRecord)

		et.bot.Send(&TelegramChat, sb)
		for _, game := range response.GameResult.Games {
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
