package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func (et *etubot) settleAll(isAuto bool) {
	games, err := et.activeLoots()
	if err != nil {
		et.sendError(fmt.Errorf("error fetching active loots: %v", err))
		return
	}

	totalSettled := 0
	for _, game := range games {
		if game.canSettle() {
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

func (et *etubot) raid() {
	teams, err := et.allTeams()
	if err != nil {
		et.sendError(fmt.Errorf("error finding teams team:%v", err))
		return
	}

	queue := 0
	for _, team := range teams {

		if !et.teamIsAvailable(team.ID) {
			// lg := fmt.Sprintf("Cannot attack, team %d is not available", team.ID)
			// log.Info(lg)
			// et.bot.Send(TelegramChat, lg)
			continue
		}

		et.pollGamesAndAttack(team)
		queue++
	}

	if !et.isAuto && queue == 0 {
		et.bot.Send(TelegramChat, "All teams are busy.")
	}
}

func (et *etubot) activeLoots() ([]Game, error) {
	var games []Game
	for i := 0; i < len(wallets); i++ {

		var response GamesResponse
		err := makeRequest(fmt.Sprintf(LootURL, wallets[i]), &response)
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

func (et *etubot) allTeams() ([]*Team, error) {
	var teams []*Team
	for i := 0; i < len(wallets); i++ {
		var response TeamsResponse
		err := makeRequest(fmt.Sprintf(TeamsURL, wallets[i]), &response)
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

func (et *etubot) pollGamesAndAttack(team *Team) {
	errorCount := 0
	lastBlock := uint64(0)
	et.bot.Send(TelegramChat, fmt.Sprintf("Finding game using team #%d", team.ID))

	for {
		bestBlock, err := et.client.BlockNumber(context.Background())
		if err != nil {
			et.sendError(fmt.Errorf("error getting block number: %v", err))
			continue
		}

		if bestBlock == lastBlock {
			continue
		}

		filterOpts := &bind.FilterOpts{Context: context.Background(), Start: bestBlock}
		gamesIter, err := et.idleContract.FilterStartGame(filterOpts)
		if err != nil {
			et.sendError(fmt.Errorf("error filtering start game: %v", err))
			continue
		}

		var target *Game
		for gamesIter.Next() {
			gameInfo, err := et.idleContract.GetGameBasicInfo(&bind.CallOpts{Context: context.Background()}, gamesIter.Event.GameId)
			if err != nil {
				et.sendError(fmt.Errorf("error getting game info: %v", err))
				continue
			}

			teamInfo, err := et.idleContract.GetTeamInfo(&bind.CallOpts{Context: context.Background()}, gamesIter.Event.TeamId)
			if err != nil {
				et.sendError(fmt.Errorf("error getting team info: %v", err))
				continue
			}

			gameAge := time.Since(time.Unix(int64(gameInfo.StartTime), 0))
			strengthDiff := team.Strength - int(teamInfo.BattlePoint)
			if strengthDiff >= 20 && gameAge < (3*time.Second) {
				target = &Game{ID: gamesIter.Event.GameId.Int64(), DefensePoint: int(teamInfo.BattlePoint), StartTime: int64(gameInfo.StartTime)}
				break
			}
		}

		if target == nil {
			continue
		}

		log.Infof("Game: %d, opponent strength: %d, team strength: %d, start time:%s", target.ID, target.DefensePoint, team.Strength, time.Since(time.Unix(target.StartTime, 0)).Truncate(time.Second))

		err = et.attack(target, team)
		if err != nil {
			errorCount++
			if errorCount >= 3 {
				et.bot.Send(TelegramChat, fmt.Sprintf("More than 3 errors while trying to attack using team %d. %v", team.ID, err))
				return
			}

			log.Error("error attacking:", err)
			continue
		}

		break
	}
}

func (et *etubot) attack(game *Game, team *Team) error {

	strengthDiff := team.Strength - game.DefensePoint
	lg := fmt.Sprintf("Attacking %d using %d, strength advantage: %d.", game.ID, team.ID, strengthDiff)
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
			et.bot.Send(TelegramChat, fmt.Sprintf("Game #%d attack successful by team #%d, defense adv: %d.\nhttps://snowtrace.io/tx/%s", game.ID, team.ID, strengthDiff, tx.Hash().String()), MsgSendOptions)
			return nil
		}

		log.Info("Attack failed, retrying")
		return fmt.Errorf("transaction failed on: %d", game.ID)
	}
}
