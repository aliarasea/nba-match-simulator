package model

type Player struct {
	Id           int32   `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Position     string  `json:"position"`
	HeightFeet   float32 `json:"height_feet"`
	HeightInches float32 `json:"height_inches"`
	WeightPounds float32 `json:"weight_pounds"`
	Team         Team    `json:"team"`
	Stats        Stats   `json:"stats"`
}

type PlayersData struct {
	Player []Player `json:"data"`
	Meta   Meta     `json:"meta"`
}
