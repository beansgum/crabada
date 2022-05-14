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
		return fmt.Errorf("error fetching all teams: %v", err)
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

		activeGameInfo, err := et.crabCaller.GetActiveGameInfo(callOpts, big.NewInt(gameID))
		if err != nil {
			return nil, fmt.Errorf("error fetching active game info: %v", err)
		}

		battleInfo, err := et.idleContract.GetGameBattleInfo(callOpts, big.NewInt(gameID))
		if err != nil {
			return nil, fmt.Errorf("error fetching game battle info by id: %v", err)
		}

		game := Game{
			ID:              gameID,
			StartTime:       int64(activeGameInfo.StartTime),
			DefensePoint:    int(activeGameInfo.DefBattlePoint),
			AttackPoint:     int(activeGameInfo.AtkBattlePoint),
			AttackTeamID:    battleInfo.AttackTeamId.Int64(),
			AttackTeamOwner: activeGameInfo.AtkOwner.String(),
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

	log.Infof("Crab %d, lockto: %d, time since locked: %s status: %d", crabID, crab.LockTo.Int64(), time.Since(time.Unix(crab.LockTo.Int64(), 0)), crab.Status.Int64())
	return time.Since(time.Unix(crab.LockTo.Int64(), 0)) > 0 && crab.Status.Int64() == 0
}

func (et *etubot) allTeams() ([]*Team, error) {

	teamIDs := []int64{1, 2, 3, 4} // hardcoded team ids
	callOpts := &bind.CallOpts{Context: context.Background()}
	var teams []*Team

	teamInfos, err := et.crabCaller.GetTeamInfos(callOpts, []*big.Int{big.NewInt(1),
		big.NewInt(2), big.NewInt(3), big.NewInt(4)}) // hardcoded team ids
	if err != nil {
		return nil, fmt.Errorf("error fetching team infos: %v", err)
	}
	for i, id := range teamIDs {
		status := "LOCK"
		if time.Since(time.Unix(teamInfos.LockTos[i].Int64(), 0)) > 0 && teamInfos.CurrentGameIds[i].Int64() == 0 {
			status = "AVAILABLE"
		}

		team := &Team{
			ID:            id,
			Strength:      int(teamInfos.BattlePoints[i]),
			Wallet:        teamInfos.TeamOwners[i].String(),
			Status:        status,
			CurrentGameID: teamInfos.CurrentGameIds[i].Int64(),
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

	return team.Status == "AVAILABLE"
}

func (et *etubot) findMyLootTeam(gameID int64) (*Team, error) {

	callOpts := &bind.CallOpts{Context: context.Background()}

	myTeamInfo, err := et.crabCaller.GetAttackTeam(callOpts, big.NewInt(gameID))
	if err != nil {
		return nil, fmt.Errorf("error fetching game attack team info: %v", err)
	}

	// TODO: verify team

	return &Team{
		ID:       myTeamInfo.AttackTeamId.Int64(),
		Strength: int(myTeamInfo.BattlePoint),
		Wallet:   myTeamInfo.TeamOwner.String()}, nil
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
				et.bot.Send(TelegramChat, fmt.Sprintf("Game #%d Team #%d has been settled, but did not get confirmed in 1 minute.\nhttps://snowtrace.io/tx/%s", gameID, team.ID, tx.Hash().String()), MsgSendOptions)
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

			log.Info("Game needs reinforcement:", game.ID)

			attackTeam := strings.ToLower(game.AttackTeamOwner)
			attackTeamCrabs := reinforcementCrabs[attackTeam]

			strengthNeeded, err := game.requiredReinforceStrength(et)
			if err != nil {
				log.Info("Error getting strength needed")
				continue
			}

			if strengthNeeded <= 0 {
				log.Info("Game got invalid needed strength:", game.ID, strengthNeeded)
				continue
			}

			if strengthNeeded > 220 { // hardcoded strength filtering
				continue
			}

			for _, crabID := range attackTeamCrabs {
				if !et.crabIsAvailable(crabID) {
					continue
				}

				// hardcoded strength filtering, sorry
				if strengthNeeded > 206 && crabID == 2584 {
					continue
				} else if strengthNeeded > 218 && crabID == 4278 {
					continue
				} else if strengthNeeded > 209 && crabID == 7272 {
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

				waitStart := time.Now()
				for {
					receipt, err := et.client.TransactionReceipt(context.Background(), tx.Hash())
					if err != nil {
						if err != ethereum.NotFound {
							log.Error("error:", err)
						}
						time.Sleep(5 * time.Second)
						if time.Since(waitStart) > 2*time.Minute {
							log.Errorf("transaction failed when reinforcing game: %d, did not get receipt after 2 minutes", game.ID)
							return
						}
						continue
					}

					if receipt.Status == types.ReceiptStatusSuccessful {
						txt := fmt.Sprintf("Game #%d reinforced with strength 220, battle points: %d vs %d.\nhttps://snowtrace.io/tx/%s", game.ID, 220+game.AttackPoint, game.DefensePoint, tx.Hash())
						et.bot.Send(TelegramChat, txt, MsgSendOptions)
					} else {
						log.Info("reinforce failed failed, retrying")
					}

					break
				}

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
	timeoutCount := 0
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
		callOpts := &bind.CallOpts{Context: context.Background()}

		txAuth, err := et.txAuth(team.Wallet, true)
		if err != nil {
			log.Error(err)
			errorCount++
			if errorCount >= 2 {
				et.bot.Send(TelegramChat, fmt.Sprintf("2 errors while trying to attack using team %d. %v", team.ID, err))
				return
			}
			continue
		}

		var startGameEvent *idlegame.IdlegameStartGame
		select {
		case event := <-channel:
			startGameEvent = event
		case <-time.After(2 * time.Minute):
			log.Info("Timeout reading from channel")
			timeoutCount++
			if timeoutCount >= 5 {
				log.Info("%d consecutive timeout, reconnect api")
				et.bot.Send(TelegramChat, fmt.Sprintf("%d consecutive timeout, reconnect api", timeoutCount))
				et.reconnect()
				return
			}
			continue
		}

		if !strings.EqualFold(startGameEvent.Raw.Address.String(), IdleContractAddress) {
			log.Info("Ignoring game non crabda contract:", startGameEvent.Raw.Address.String())
			continue
		}

		gameInfo, err := et.crabCaller.GetGameDefTeamInfo(callOpts, startGameEvent.GameId)
		if err != nil {
			et.sendError(fmt.Errorf("error getting game info: %v", err))
			return
		}

		gameAge := time.Since(time.Unix(int64(gameInfo.StartTime), 0))
		strengthDiff := team.Strength - int(gameInfo.BattlePoint)
		// Only attempting attacks on < 3s old games with >10 strength advantage
		if strengthDiff >= 10 && gameAge < (3*time.Second) {
			targetGame := &Game{ID: startGameEvent.GameId.Int64(), DefensePoint: int(gameInfo.BattlePoint)}
			go log.Infof("Game: %d, opponent strength: %d, team strength: %d, start time:%s", targetGame.ID, targetGame.DefensePoint, team.Strength, gameAge.Truncate(time.Second))

			battleInfo, err := et.idleContract.GetGameBattleInfo(callOpts, big.NewInt(targetGame.ID))
			if err != nil {
				log.Errorf("error getting defense team info 2: %v", err)
				continue
			}

			if battleInfo.AttackTeamId.Int64() != 0 {
				log.Error("error game looted:", targetGame.ID)
				continue
			}

			err = et.attack(txAuth, targetGame, team, strengthDiff)
			if err != nil {
				errorCount++
				if errorCount >= 2 {
					et.bot.Send(TelegramChat, fmt.Sprintf("2 errors while trying to attack using team %d. %v", team.ID, err))
					return
				}

				log.Error("error attacking:", err)
				continue
			}

			et.gamesMu.Lock()
			et.games[targetGame.ID] = 1
			et.gamesMu.Unlock()
			break
		} else {
			// log.Infof("Cant attack, strenght diff %d, age %s", strengthDiff, gameAge)
		}
	}
}

func (et *etubot) attack(auth *bind.TransactOpts, game *Game, team *Team, strengthDiff int) error {

	go log.Infof("Attacking %d using %d, strength advantage: %d.", game.ID, team.ID, strengthDiff)

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
