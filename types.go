package main

type Team struct {
	ID       int
	Strength int
	Wallet   string
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
	ID              int           `json:"game_id"`
	StartTime       int64         `json:"start_time"`
	EndTime         int64         `json:"end_time"`
	DefensePoint    int           `json:"defense_point"`
	AttackPoint     int           `json:"attack_point"`
	AttackTeamID    int           `json:"attack_team_id"`
	AttackTeamOwner string        `json:"attack_team_owner"`
	WinnerTeamID    string        `json:"winner_team_id"`
	Process         []GameProcess `json:"process"`
}

type GameProcess struct {
	Action string `json:"action"`
}
