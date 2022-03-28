package player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nway2inside/model"
	"os"
)

func getPlayersData(page int) model.PlayersData {
	if page <= 0 {
		page = 1
	}

	response, err := http.Get(fmt.Sprintf("https://www.balldontlie.io/api/v1/players?page=%d", page))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	var playersData model.PlayersData

	err = json.Unmarshal(responseData, &playersData)

	return playersData
}

func GetAllPlayers() []model.Player {
	var playersData model.PlayersData

	playersData = getPlayersData(1)

	totalPages := playersData.Meta.TotalPages

	players := playersData.Player

	for i := 2; i <= totalPages; i++ {
		playersData = getPlayersData(i)
		if playersData.Player != nil {
			players = append(players, playersData.Player...)
		}
	}
	return players
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	players := GetAllPlayers()
	err := json.NewEncoder(w).Encode(players)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
