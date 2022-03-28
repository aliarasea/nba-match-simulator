package model

type Team struct {
	Id           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
	City         string `json:"city"`
	Conference   string `json:"conference"`
	Division     string `json:"division"`
	FullName     string `json:"full_name"`
	Name         string `json:"name"`
}

type TeamsData struct {
	Teams []Team `json:"data"`
	Meta  Meta   `json:"meta"`
}
