package team

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"nway2inside/model"
	"os"
)

func GetTeamsData(page int) model.TeamsData {
	if page <= 0 {
		page = 1
	}

	response, err := http.Get(fmt.Sprintf("https://www.balldontlie.io/api/v1/teams?page=%d", page))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	var teamsData model.TeamsData

	err = json.Unmarshal(responseData, &teamsData)

	return teamsData
}

func GetAllTeams() []model.Team {
	var teamsData model.TeamsData

	teamsData = GetTeamsData(1)

	totalPages := teamsData.Meta.TotalPages

	teams := teamsData.Teams

	for i := 2; i <= totalPages; i++ {
		teamsData = GetTeamsData(i)
		if teamsData.Teams != nil {
			teams = append(teams, teamsData.Teams...)
		}
	}
	return teams
}

func GetTeams(w http.ResponseWriter, r *http.Request) {
	teams := GetAllTeams()
	err := json.NewEncoder(w).Encode(teams)
	if err != nil {
		log.Fatal(err)
	}
}
