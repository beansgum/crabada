package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/c-ollins/crabada/idlegame"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

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

func (et *etubot) crabIsAvailable(crabID int64) bool {
	crab, err := et.crabForID(crabID)
	if err != nil {
		et.sendError(fmt.Errorf("error fetching crab: %v", err))
		return false
	}

	if crab.Status == "AVAILABLE" {
		return true
	}

	return false
}

func (et *etubot) crabForID(crabID int64) (*Crab, error) {
	var response CrabResponse
	err := makeRequest(fmt.Sprintf("https://idle-api.crabada.com/public/idle/crabada/info/%d", crabID), &response)
	if err != nil {
		return nil, fmt.Errorf("error fetching game by id: %v", err)
	}

	if response.ErrorCode != "" {
		return nil, fmt.Errorf("error fetching game by id: %s, message: %s", response.ErrorCode, response.Message)
	}

	return &response.Crab, nil
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

	auth, err := et.txAuth(team.Wallet, false)
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

func (et *etubot) reinforceAttacks() {
	games, err := et.activeLoots()
	if err != nil {
		et.sendError(fmt.Errorf("error fetching active games:%v", err))
		return
	}

	for _, game := range games {
		if game.needsReinforcement() {

			var crabsResponse CrabsResponse
			err = makeRequest(fmt.Sprintf(CrabsURL, game.AttackTeamOwner), &crabsResponse)
			if err != nil {
				log.Error("Error fetching available crabs:", err)
				continue
			}

			if crabsResponse.ErrorCode != "" {
				log.Infof("error fetching crabs: %s, message: %s", crabsResponse.ErrorCode, crabsResponse.Message)
				continue
			}

			strengthNeeded := game.DefensePoint - game.AttackPoint
			var chosenCrab *Crab
			for _, crab := range crabsResponse.Data.Crabs {
				if crab.BattlePoint >= strengthNeeded && (chosenCrab == nil || chosenCrab.BattlePoint > crab.BattlePoint) {
					if et.crabIsAvailable(int64(crab.CrabadaID)) {
						chosenCrab = &crab
					} else {
						log.Infof("Tried to pick crab %d for game %d but not available", crab.CrabadaID, game.ID)
					}
				}
			}

			if chosenCrab != nil {
				log.Infof("Reinforcing game #%d using crab #%d", game.ID, chosenCrab.ID)
				auth, err := et.txAuth(game.AttackTeamOwner, false)
				if err != nil {
					et.sendError(fmt.Errorf("error creating tx auth: %v", err))
					return
				}

				tx, err := et.idleContract.ReinforceAttack(auth, big.NewInt(game.ID), big.NewInt(int64(chosenCrab.CrabadaID)), big.NewInt(0))
				if err != nil {
					et.sendError(fmt.Errorf("error reinforcing defense: %v", err))
					return
				}

				txt := fmt.Sprintf("Game #%d reinforced with strength %d, battle points: %d vs %d.\nhttps://snowtrace.io/tx/%s", game.ID, chosenCrab.BattlePoint, chosenCrab.BattlePoint+game.AttackPoint, game.DefensePoint, tx.Hash())
				et.bot.Send(TelegramChat, txt, MsgSendOptions)
			} else {
				log.Infof("Cannot reinforce game #%d, no available crab", game.ID)
			}
		}
	}
}

func (et *etubot) raid() {
	teams, err := et.allTeams()
	if err != nil {
		et.sendError(fmt.Errorf("error finding all team:%v", err))
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

func (et *etubot) pollGamesAndAttack(team *Team) {
	errorCount := 0
	et.bot.Send(TelegramChat, fmt.Sprintf("Finding game using team #%d", team.ID))

	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	channel := make(chan *idlegame.IdlegameStartGame)

	log.Info("Subscribing to contract")
	sub, err := et.idleContract.WatchStartGame(watchOpts, channel)
	if err != nil {
		et.sendError(fmt.Errorf("error watching start game: %v", err))
		return
	}
	defer sub.Unsubscribe()

	for {
		startGameEvent := <-channel

		callOpts := &bind.CallOpts{Context: context.Background()}

		gameInfo, err := et.idleContract.GetGameBasicInfo(callOpts, startGameEvent.GameId)
		if err != nil {
			et.sendError(fmt.Errorf("error getting game info: %v", err))
			return
		}

		teamInfo, err := et.idleContract.GetTeamInfo(callOpts, startGameEvent.TeamId)
		if err != nil {
			et.sendError(fmt.Errorf("error getting game info: %v", err))
			return
		}

		gameAge := time.Since(time.Unix(int64(gameInfo.StartTime), 0))
		strengthDiff := team.Strength - int(teamInfo.BattlePoint)
		if strengthDiff >= 20 && gameAge < (3*time.Second) {
			targetGame := &Game{ID: startGameEvent.GameId.Int64(), DefensePoint: int(teamInfo.BattlePoint), StartTime: int64(gameInfo.StartTime)}
			go log.Infof("Game: %d, opponent strength: %d, team strength: %d, start time:%s", targetGame.ID, targetGame.DefensePoint, team.Strength, time.Since(time.Unix(targetGame.StartTime, 0)).Truncate(time.Second))

			err = et.attack(targetGame, team)
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
}

func (et *etubot) attack(game *Game, team *Team) error {

	strengthDiff := team.Strength - game.DefensePoint
	go log.Infof("Attacking %d using %d, strength advantage: %d.", game.ID, team.ID, strengthDiff)
	auth, err := et.txAuth(team.Wallet, true)
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
