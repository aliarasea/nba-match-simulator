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

func GetTeams(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://www.balldontlie.io/service/v1/teams")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	var teamsData model.TeamsData

	err = json.Unmarshal(responseData, &teamsData)

	err = json.NewEncoder(w).Encode(teamsData.Teams)

	for _, team := range teamsData.Teams {
		_, err = fmt.Fprintln(w, team.Name)
	}

	if err != nil {
		log.Fatal(err)
	}
}
