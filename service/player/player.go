package player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nway2inside/helper"
	"nway2inside/model"
	"nway2inside/service/team"
	"os"
)

type Players []model.Player

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

func GetAllPlayers() Players {
	var playersData model.PlayersData

	playersData = getPlayersData(1)

	totalPages := playersData.Meta.TotalPages

	var players Players

	players = playersData.Player

	for i := 2; i <= totalPages; i++ {
		playersData = getPlayersData(i)
		if playersData.Player != nil {
			players = append(players, playersData.Player...)
		}
	}
	return players.Shuffle().AssignToTeams()
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	players := GetAllPlayers()
	err := json.NewEncoder(w).Encode(players)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

func (players Players) Shuffle() Players {
	return helper.Shuffle(players)
}

func (players Players) AssignToTeams() Players {
	teams := team.GetAllTeams()
	for _, teamItem := range teams {
		teams.Shuffle()
		for _, player := range players {
			player.Team = teamItem
		}
	}
	return players
}
