package model

import "time"

type Match struct {
	Id            int64     `json:"id"`
	HomeTeam      Team      `json:"home_team"`
	AwayTeam      Team      `json:"away_team"`
	HomeTeamScore int16     `json:"home_team_score"`
	AwayTeamScore int16     `json:"away_team_score"`
	Description   string    `json:"description"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	Status        string    `json:"status"`
	Result        string    `json:"result"`
}
