package main

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type TeamsResponse struct {
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
	TeamResult `json:"result"`
}

type TeamResult struct {
	TotalRecord int     `json:"totalRecord"`
	Teams       []*Team `json:"data"`
}

type Team struct {
	ID            int64  `json:"team_id"`
	Strength      int    `json:"battle_point"`
	Wallet        string `json:"owner"`
	Status        string `json:"status"` //AVAILABLE/LOCK
	CurrentGameID int64
}

type GamesResponse struct {
	ErrorCode   string `json:"error_code"`
	Message     string `json:"message"`
	GamesResult `json:"result"`
}

type GameResponse struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
	Game      `json:"result"`
}

type GamesResult struct {
	TotalRecord int    `json:"totalRecord"`
	Games       []Game `json:"data"`
}

type Game struct {
	ID              int64  `json:"game_id"`
	StartTime       int64  `json:"start_time"`
	DefensePoint    int    `json:"defense_point"`
	AttackPoint     int    `json:"attack_point"`
	AttackTeamID    int64  `json:"attack_team_id"`
	AttackTeamOwner string `json:"attack_team_owner"`
	BattleInfo      *BattleInfo
	Process         []GameProcess `json:"process"`
}

type BattleInfo struct {
	AttackTeamId   *big.Int
	AttackTime     uint32
	LastAttackTime uint32
	LastDefTime    uint32
	AttackId1      *big.Int
	AttackId2      *big.Int
	DefId1         *big.Int
	DefId2         *big.Int
}

func (g Game) startTime() time.Time {
	return time.Unix(g.StartTime, 0)
}

func (g Game) settleTime() time.Time {
	return g.startTime().Add(time.Duration(1) * time.Hour)
}

func (g Game) isWonOrLost() bool {
	// true if  the enemy did not reinforce or we did not reattack a reinforcement in time
	return time.Since(g.lastProcessTime()) > processIntervals || g.reinforceAttackCount() == 2
}

func (g Game) canSettle() bool {
	return g.isWonOrLost() && time.Now().After(g.settleTime())
}

func (g Game) lastProcessTime() time.Time {
	if g.reinforceAttackCount() > g.reinforceDefenseCount() {
		return g.lastAttackTime()
	}

	return g.lastDefenseTime()
}

func (g Game) requiredReinforceStrength(et *etubot) (int64, error) {

	statsInfo, err := et.idleContract.GetLootingStatsInfo(&bind.CallOpts{Context: context.Background()}, big.NewInt(g.ID))
	if err != nil {
		return 0, err
	}

	return int64(statsInfo.DefensePoint) - int64(statsInfo.AttackPoint), nil
}

func (g Game) reinforceAttackCount() int {
	count := 0
	if g.BattleInfo.AttackId1.Int64() != 0 {
		count++
	}

	if g.BattleInfo.AttackId2.Int64() != 0 {
		count++
	}

	return count
}

func (g Game) lastAttackTime() time.Time {
	return time.Unix(int64(g.BattleInfo.LastAttackTime), 0)
}

func (g Game) reinforceDefenseCount() int {
	count := 0
	if g.BattleInfo.DefId1.Int64() != 0 {
		count++
	}

	if g.BattleInfo.DefId2.Int64() != 0 {
		count++
	}

	return count
}

func (g Game) lastDefenseTime() time.Time {
	return time.Unix(int64(g.BattleInfo.LastDefTime), 0)
}

func (g Game) needsReinforcement() bool {
	return g.lastDefenseTime().After(g.lastAttackTime()) && time.Since(g.lastDefenseTime()) < processIntervals
}

type GameProcess struct {
	Action string `json:"action"`
	TxTime int64  `json:"transaction_time"`
}

type Crab struct {
	Owner  common.Address
	LockTo *big.Int
	Status *big.Int
}

type GasResponse struct {
	ErrorCode int     `json:"error_code"`
	Data      GasData `json:"data"`
}

type GasData struct {
	Fast GasPrice `json:"fast"`
}

type GasPrice struct {
	Price float64 `json:"price"`
}
