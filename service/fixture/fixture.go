package fixture

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"nway2inside/model"
	"nway2inside/service/team"
	"os"
	"time"
)

type Fixture model.Fixture

func (Fixture) CreateFixture() Fixture {
	teams := team.GetAllTeams()
	matchCount := len(teams) / 2
	var matches []model.Match
	startTime := time.Now().Local().AddDate(0, 0, 2)
	endTime := time.Now().Local().AddDate(0, 0, 2).Add(time.Minute * 48)
	for i := 0; i < matchCount; i++ {
		var match = model.Match{
			Id:          int64(i),
			HomeTeam:    teams[i],
			AwayTeam:    teams[matchCount-i],
			Description: fmt.Sprintf("%s vs %s", teams[i].Name, teams[matchCount-i].Name),
			StartTime:   startTime,
			EndTime:     endTime,
			Status:      "PENDING",
		}
		matches = append(matches, match)
	}

	return Fixture{
		Id:          int64(rand.Int()),
		Name:        "Test Fixture",
		Description: "Test fixture",
		Matches:     matches,
	}
}

func GetFixture(w http.ResponseWriter, r *http.Request) {
	var fixture Fixture
	err := json.NewEncoder(w).Encode(fixture.CreateFixture())
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
