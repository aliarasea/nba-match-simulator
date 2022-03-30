package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"nway2inside/service/fixture"
	"nway2inside/service/player"
	"nway2inside/service/team"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the NBA simulator!")
	if err != nil {
		fmt.Println("Something went wrong!")
		return
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	_, err := fmt.Fprintf(w, string(reqBody))
	if err != nil {
		fmt.Println("Something went wrong!")
		return
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api", api).Methods("POST")
	myRouter.HandleFunc("/api/teams", team.GetTeams).Methods("GET")
	myRouter.HandleFunc("/api/players", player.GetPlayers).Methods("GET")
	myRouter.HandleFunc("/api/fixture", fixture.GetFixture).Methods("GET")
	log.Fatal(http.ListenAndServe(":7676", myRouter))
}

func main() {
	handleRequests()
}
