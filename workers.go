package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"
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

func (et *etubot) watchRaidGas() {

	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	channel := make(chan *idlegame.IdlegameFight)

	sub, err := et.idleContract.WatchFight(watchOpts, channel)
	if err != nil {
		et.sendError(fmt.Errorf("error watching fights: %v", err))
		return
	}

	defer sub.Unsubscribe()

	log.Info("Watching raid gas")
	var lastHighestFee = big.NewInt(0)
	for {
		event := <-channel
		if event.Turn.Int64() != 0 {
			continue
		}
		tx, _, err := et.client.TransactionByHash(context.Background(), event.Raw.TxHash)
		if err != nil {
			log.Errorf("error getting fight event tx: %v", err)
			continue
		}

		if tx.GasTipCap().Int64() <= 0 {
			continue
		}

		if !strings.EqualFold(tx.To().String(), IdleContractAddress) {
			log.Info("Ignoring non crabda contract:", tx.To())
			continue
		}

		receipt, err := et.client.TransactionReceipt(context.Background(), event.Raw.TxHash)
		if err != nil {
			log.Errorf("error getting fight event receipt: %v", err)
			continue
		}

		header, err := et.client.HeaderByHash(context.Background(), receipt.BlockHash)
		if err != nil {
			log.Errorf("error getting fight block header: %v", err)
			continue
		}

		gasPrice := big.NewInt(0).Add(header.BaseFee, tx.GasTipCap())
		if gasPrice.Cmp(tx.GasFeeCap()) == 1 {
			gasPrice = tx.GasFeeCap()
		}

		if gasPrice.Cmp(lastHighestFee) == 1 {
			lastHighestFee = gasPrice
		}

		et.raidGasMu.Lock()
		if time.Since(et.raidLastUpdate) > 60*time.Second || et.raidGasPrice.Cmp(lastHighestFee) == -1 {
			et.raidGasPrice = lastHighestFee
			et.raidLastUpdate = time.Now()
			log.Info("Raid gas now:", ToGwei(lastHighestFee))
			lastHighestFee = big.NewInt(0)
		}
		et.raidGasMu.Unlock()

	}
}

func (et *etubot) auto() {
	et.bot.Send(TelegramChat, "Etubot running on auto.")
	for {
		log.Info("Auto running")
		// settle ready games

		et.gasMu.RLock()
		gasPrice := et.gasPrice
		et.gasMu.RUnlock()

		if gasPrice.Cmp(GasLimit) >= 0 {
			log.Info("Gas is too high")
			// notify high gas
			// continue
		} else {
			et.settleAll(true)
			et.reinforceAttacks()
		}

		et.raidGasMu.RLock()
		raidGas := big.NewInt(0).Add(et.raidGasPrice, RaidGasExtra)
		et.raidGasMu.RUnlock()

		if !(raidGas.Cmp(RaidGasLimit) >= 0) && et.lootingActive {
			et.raid()
		}

		time.Sleep(30 * time.Second)
	}
}
