package main

import "time"

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
	ID       int64  `json:"team_id"`
	Strength int    `json:"battle_point"`
	Wallet   string `json:"owner"`
	Status   string `json:"status"` //AVAILABLE/LOCK
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
	ID              int64         `json:"game_id"`
	StartTime       int64         `json:"start_time"`
	EndTime         int64         `json:"end_time"`
	DefensePoint    int           `json:"defense_point"`
	AttackPoint     int           `json:"attack_point"`
	AttackTeamID    int64         `json:"attack_team_id"`
	AttackTeamOwner string        `json:"attack_team_owner"`
	WinnerTeamID    string        `json:"winner_team_id"`
	Process         []GameProcess `json:"process"`
}

func (g Game) startTime() time.Time {
	return time.Unix(g.StartTime, 0)
}

func (g Game) settleTime() time.Time {
	return g.startTime().Add(time.Duration(1) * time.Hour)
}

func (g Game) isWonOrLost() bool {
	latestProcess := g.lastProcess()

	// true if  the enemy did not reinforce or we did not reattack a reinforcement in time
	return time.Since(latestProcess.txTime()) > processIntervals
}

func (g Game) canSettle() bool {
	return g.isWonOrLost() && time.Now().After(g.settleTime())
}

func (g Game) lastProcess() GameProcess {
	return g.Process[len(g.Process)-1]
}

func (g Game) needsReinforcement() bool {
	latestProcess := g.lastProcess()
	if latestProcess.Action == actionReinforceDefense {
		if time.Since(latestProcess.txTime()) < processIntervals {
			return true
		}
	}

	return false
}

type GameProcess struct {
	Action string `json:"action"`
	TxTime int64  `json:"transaction_time"`
}

func (gp GameProcess) txTime() time.Time {
	return time.Unix(gp.TxTime, 0)
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
