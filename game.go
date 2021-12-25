package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/c-ollins/crabada/idlegame"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func (et *etubot) addActiveGames() error {
	teams, err := et.allTeams()
	if err != nil {
		return err
	}

	et.gamesMu.Lock()
	defer et.gamesMu.Unlock()
	for _, team := range teams {
		if team.CurrentGameID != 0 {
			et.games[team.CurrentGameID] = 1
		}
	}

	return nil
}

func (et *etubot) activeLoots() ([]Game, error) {
	var games []Game
	et.gamesMu.Lock()
	defer et.gamesMu.Unlock()
	for gameID, _ := range et.games {
		callOpts := &bind.CallOpts{Context: context.Background()}

		gameInfo, err := et.idleContract.GetGameBasicInfo(callOpts, big.NewInt(gameID))
		if err != nil {
			return nil, fmt.Errorf("error fetching game battle info by id: %v", err)
		}

		oppTeamInfo, err := et.idleContract.GetTeamInfo(&bind.CallOpts{Context: context.Background()}, gameInfo.TeamId)
		if err != nil {
			return nil, fmt.Errorf("error fetching game team info by id: %v", err)
		}

		battleInfo, err := et.idleContract.GetGameBattleInfo(callOpts, big.NewInt(gameID))
		if err != nil {
			return nil, fmt.Errorf("error fetching game battle info by id: %v", err)
		}

		myTeamInfo, err := et.idleContract.GetTeamInfo(&bind.CallOpts{Context: context.Background()}, battleInfo.AttackTeamId)
		if err != nil {
			return nil, fmt.Errorf("error fetching game team info by id: %v", err)
		}

		// log.Infof("Attack time: %d, last attack time: %d, last defense time %d", battleInfo.AttackTime, battleInfo.LastAttackTime, battleInfo.LastDefTime, battleInfo.LastDefTime > battleInfo.LastAttackTime)
		// log.Infof("AttackID %d %d, Def Id %d, %d", battleInfo.AttackId1, battleInfo.AttackId2, battleInfo.DefId1.Int64(), battleInfo.DefId2.Int64())
		game := Game{
			ID:              gameID,
			StartTime:       int64(gameInfo.StartTime),
			DefensePoint:    int(oppTeamInfo.BattlePoint),
			AttackPoint:     int(myTeamInfo.BattlePoint),
			AttackTeamID:    battleInfo.AttackTeamId.Int64(),
			AttackTeamOwner: myTeamInfo.Owner.String(),
			BattleInfo:      (*BattleInfo)(&battleInfo),
		}

		games = append(games, game)
	}

	return games, nil
}

func (et *etubot) crabForID(crabID int64) (*Crab, error) {
	callOpts := &bind.CallOpts{Context: context.Background()}

	crab, err := et.idleContract.CrabaInfos(callOpts, big.NewInt(crabID))
	if err != nil {
		return nil, err
	}

	return (*Crab)(&crab), nil
}

func (et *etubot) crabIsAvailable(crabID int64) bool {
	crab, err := et.crabForID(crabID)
	if err != nil {
		log.Errorf("Error fetching crabada info:", err)
		return false
	}

	// log.Infof("Crab %d, lockto: %d, status: %d", crabID, crab.LockTo.Int64(), crab.Status.Int64())
	return time.Since(time.Unix(crab.LockTo.Int64(), 0)) > 0 && crab.Status.Int64() == 0
}

func (et *etubot) allTeams() ([]*Team, error) {

	teamIDs := []int64{2411, 2290, 2279, 2463, 2462, 2461}
	callOpts := &bind.CallOpts{Context: context.Background()}
	var teams []*Team
	for _, id := range teamIDs {
		teamInfo, err := et.idleContract.GetTeamInfo(callOpts, big.NewInt(id))
		if err != nil {
			log.Error("Error fetching team info:", err)
			continue
		}

		status := "LOCK"
		if time.Since(time.Unix(teamInfo.LockTo.Int64(), 0)) > 0 && teamInfo.CurrentGameId.Int64() == 0 {
			status = "AVAILABLE"
		}

		team := &Team{
			ID:            id,
			Strength:      int(teamInfo.BattlePoint),
			Wallet:        teamInfo.Owner.String(),
			Status:        status,
			CurrentGameID: teamInfo.CurrentGameId.Int64(),
		}

		teams = append(teams, team)
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

	callOpts := &bind.CallOpts{Context: context.Background()}

	battleInfo, err := et.idleContract.GetGameBattleInfo(callOpts, big.NewInt(gameID))
	if err != nil {
		return nil, fmt.Errorf("error fetching game battle info by id: %v", err)
	}

	myTeamInfo, err := et.idleContract.GetTeamInfo(callOpts, battleInfo.AttackTeamId)
	if err != nil {
		return nil, fmt.Errorf("error fetching game team info by id: %v", err)
	}

	// TODO: verify team

	return &Team{
		ID:       battleInfo.AttackTeamId.Int64(),
		Strength: int(myTeamInfo.BattlePoint),
		Wallet:   myTeamInfo.Owner.String()}, nil
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
		et.gamesMu.Lock()
		delete(et.games, gameID)
		et.gamesMu.Unlock()
		break
	}
}

func (et *etubot) reinforceAttacks() {
	log.Info("reinfocing")
	games, err := et.activeLoots()
	if err != nil {
		et.sendError(fmt.Errorf("error fetching active games:%v", err))
		return
	}

	for _, game := range games {
		if game.needsReinforcement() {

			attackTeam := strings.ToLower(game.AttackTeamOwner)
			attackTeamCrabs := reinforcementCrabs[attackTeam]

			for _, crabID := range attackTeamCrabs {
				log.Info("available: ", et.crabIsAvailable(crabID))
			}

			strengthNeeded, err := game.requiredReinforceStrength(et)
			if err != nil {
				log.Info("Error getting strength needed")
				continue
			}

			if strengthNeeded > 220 {
				continue
			}

			for _, crabID := range attackTeamCrabs {
				if !et.crabIsAvailable(crabID) {
					continue
				}

				log.Infof("Reinforcing game #%d using crab #%d", game.ID, crabID)
				auth, err := et.txAuth(game.AttackTeamOwner, false)
				if err != nil {
					et.sendError(fmt.Errorf("error creating tx auth: %v", err))
					return
				}

				tx, err := et.idleContract.ReinforceAttack(auth, big.NewInt(game.ID), big.NewInt(int64(crabID)), big.NewInt(0))
				if err != nil {
					et.sendError(fmt.Errorf("error reinforcing defense: %v", err))
					return
				}

				txt := fmt.Sprintf("Game #%d reinforced with strength 220, battle points: %d vs %d.\nhttps://snowtrace.io/tx/%s", game.ID, 220+game.AttackPoint, game.DefensePoint, tx.Hash())
				et.bot.Send(TelegramChat, txt, MsgSendOptions)
				break
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

	log.Info("got teams", len(teams))

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
		log.Info("Reading to contract channel")
		var startGameEvent *idlegame.IdlegameStartGame
		select {
		case event := <-channel:
			startGameEvent = event
		case <-time.After(2 * time.Minute):
			log.Info("Timeout reading from channel")
			continue
		}

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

			et.gamesMu.Lock()
			et.games[targetGame.ID] = 1
			et.gamesMu.Unlock()
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
