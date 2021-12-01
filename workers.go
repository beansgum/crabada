package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

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
	for {
		bestBlock, err := et.client.BlockNumber(context.Background())
		if err != nil {
			et.sendError(fmt.Errorf("error getting block number: %v", err))
			return
		}

		log.Info("Bestblock:", bestBlock)
		filterOpts := &bind.FilterOpts{Context: context.Background(), Start: bestBlock}
		gamesIter, err := et.idleContract.FilterStartGame(filterOpts)
		if err != nil {
			et.sendError(fmt.Errorf("error filtering start game: %v", err))
			return
		}

		for gamesIter.Next() {
			gameInfo, err := et.idleContract.GetGameBasicInfo(&bind.CallOpts{Context: context.Background()}, gamesIter.Event.GameId)
			if err != nil {
				et.sendError(fmt.Errorf("error getting game info: %v", err))
				return
			}

			teamInfo, err := et.idleContract.GetTeamInfo(&bind.CallOpts{Context: context.Background()}, gamesIter.Event.TeamId)
			if err != nil {
				et.sendError(fmt.Errorf("error getting game info: %v", err))
				return
			}

			gameAge := time.Since(time.Unix(int64(gameInfo.StartTime), 0)).Truncate(time.Second)

			log.Info("Game:", gamesIter.Event.GameId, "Strength:", teamInfo.BattlePoint, "Start:", gameAge)
		}
	}
	// watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// channel := make(chan *idlegame.IdlegameStartGame)

	// // Start a goroutine which watches new events
	// go func() {
	// 	log.Info("Subscribing to contract")
	// 	sub, err := et.idleContract.WatchStartGame(watchOpts, channel)
	// 	if err != nil {
	// 		et.sendError(fmt.Errorf("error watching start game: %v", err))
	// 		return
	// 	}
	// 	defer sub.Unsubscribe()
	// }()

	// log.Info("Watch start game")

	// for {
	// 	event := <-channel
	// 	log.Infof("New game id: %d, team id: %d", event.GameId, event.TeamId)
	// }
}

func (et *etubot) auto() {
	et.bot.Send(TelegramChat, "Etubot running on auto.")
	for {
		log.Info("Auto running")
		// settle ready games

		et.gasMu.RLock()
		gasPrice := big.NewInt(0).Add(et.gasPrice, big.NewInt(30000000000))
		et.gasMu.RUnlock()
		limit := big.NewInt(200000000000) //200gwei

		if gasPrice.Cmp(limit) >= 0 {
			// notify high gas
			// continue
		} else {
			et.settleAll(true)
			et.raid()
		}

		time.Sleep(30 * time.Second)
	}
}
