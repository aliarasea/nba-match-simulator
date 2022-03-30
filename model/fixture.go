package model

type Fixture struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Matches     []Match `json:"matches"`
}
