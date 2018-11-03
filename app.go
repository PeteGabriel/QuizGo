package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var questions []Question

type App struct {
	Router *mux.Router
}

func seedQuestions() {
	questions = make([]Question, 3)
	questions[0] = Question{
		Content:    "What is the national animal of China?",
		corrOption: "Giant Panda",
		Options:    [4]string{"Giant Panda", "Panda", "Red Panda", "Polar Bear"},
	}
	questions[1] = Question{
		Content:    "What is the noisiest city in the world?",
		corrOption: "Hong Kong",
		Options:    [4]string{"Lisbon", "London", "New York", "Hong Kong"},
	}
	questions[2] = Question{
		Content:    "What is a very cold part of Russia?",
		corrOption: "Siberia",
		Options:    [4]string{"Magadan", "Krasnoyarsk", "Siberia", "Moscow"},
	}
}

// Run func should be used to start the server
func (a *App) Run(addr string) {
	a.initialize()
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initialize() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/quiz", getQuiz).Methods("GET")
	a.Router.HandleFunc("/quiz/results", saveResults).Methods("POST")

	seedQuestions()
}

func getQuiz(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ", r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(questions); err != nil {
		panic(err)
	}
}

func saveResults(w http.ResponseWriter, r *http.Request) {

}
