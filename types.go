package main

type Team struct {
	ID       int
	Strength int
}

type GameResponse struct {
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
	GameResult `json:"result"`
}

type GameResult struct {
	TotalRecord int    `json:"totalRecord"`
	Games       []Game `json:"data"`
}

type Game struct {
	ID              int           `json:"game_id"`
	StartTime       int64         `json:"start_time"`
	EndTime         int64         `json:"end_time"`
	DefensePoint    int           `json:"defense_point"`
	AttackTeamID    int           `json:"attack_team_id"`
	AttackTeamOwner string        `json:"attack_team_owner"`
	WinnerTeamID    string        `json:"winner_team_id"`
	Process         []GameProcess `json:"process"`
}

type GameProcess struct {
	Action string `json:"action"`
}
