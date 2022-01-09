package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/c-ollins/crabada/idlegame"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

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

func (et *etubot) watchStartGame() {
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	channel := make(chan *idlegame.IdlegameStartGame)

	log.Info("Subscribing to contract")
	sub, err := et.idleContract.WatchStartGame(watchOpts, channel)
	if err != nil {
		et.sendError(fmt.Errorf("error watching start game: %v", err))
		return
	}
	defer log.Info("unsub")
	defer sub.Unsubscribe()

	log.Info("Watch start game")
	for {
		log.Info("Reading event")
		event := <-channel
		log.Infof("New game id: %d, team id: %d", event.GameId, event.TeamId)
		gameInfo, err := et.idleContract.GetGameBasicInfo(&bind.CallOpts{Context: context.Background()}, event.GameId)
		if err != nil {
			et.sendError(fmt.Errorf("error getting game info: %v", err))
			return
		}

		teamInfo, err := et.idleContract.GetTeamInfo(&bind.CallOpts{Context: context.Background()}, event.TeamId)
		if err != nil {
			et.sendError(fmt.Errorf("error getting game info: %v", err))
			return
		}

		gameAge := time.Since(time.Unix(int64(gameInfo.StartTime), 0)).Truncate(time.Second)

		log.Info("Game:", event.GameId, "Strength:", teamInfo.BattlePoint, "Start:", gameAge)
	}
}

func (et *etubot) auto() {
	et.bot.Send(TelegramChat, "Etubot running on auto.")
	for {
		log.Info("Auto running")
		// settle ready games

		et.gasMu.RLock()
		gasPrice := et.gasPrice
		raidGas := big.NewInt(0).Add(et.gasPrice, big.NewInt(120000000000))
		et.gasMu.RUnlock()
		limit := big.NewInt(210000000000) //200gwei

		if gasPrice.Cmp(limit) >= 0 {
			// notify high gas
			// continue
		} else {
			et.settleAll(true)
			et.reinforceAttacks()
		}

		if !(raidGas.Cmp(limit) >= 0) {
			et.raid()
		}

		time.Sleep(30 * time.Second)
	}
}
