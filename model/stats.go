package model

type Stats struct {
	Points                 int `json:"points"`
	Rebounds               int `json:"rebounds"`
	Assists                int `json:"assists"`
	Steals                 int `json:"steals"`
	Blocks                 int `json:"blocks"`
	Turnovers              int `json:"turnovers"`
	Fouls                  int `json:"fouls"`
	FieldGoalsMade         int `json:"field_goals_made"`
	FieldGoalsAttempted    int `json:"field_goals_attempted"`
	ThreePointersMade      int `json:"three_pointers_made"`
	ThreePointersAttempted int `json:"three_pointers_attempted"`
	FreeThrowsMade         int `json:"free_throws_made"`
	FreeThrowsAttempted    int `json:"free_throws_attempted"`
	PersonalFouls          int `json:"personal_fouls"`
	OffensiveRebounds      int `json:"offensive_rebounds"`
	DefensiveRebounds      int `json:"defensive_rebounds"`
}
