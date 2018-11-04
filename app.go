package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	NumQuests = 3
	GET       = "GET"
	POST      = "POST"
)

var questions []Question
var rates []int
var expectedRes [NumQuests]string

type App struct {
	Router *mux.Router
}

func init() {
	rates = make([]int, 1)
	expectedRes = [NumQuests]string{"Giant Panda", "Hong Kong", "Siberia"}
	seedQuestions()
}

// Run func should be used to start the server
func (a *App) Run(addr string) {
	handler := a.initialize()
	log.Fatal(http.ListenAndServe(addr, handler))
}

func (a *App) initialize() http.Handler {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/quiz", GetQuiz).Methods(GET)
	a.Router.HandleFunc("/quiz/results", SaveResults).Methods(POST)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	handler := c.Handler(a.Router)
	return handler
}

func GetQuiz(w http.ResponseWriter, r *http.Request) {
	log.Println(GET, r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(questions); err != nil {
		panic(err)
	}
}

func SaveResults(w http.ResponseWriter, r *http.Request) {
	log.Println(POST, r.RequestURI)
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	result := Result{}
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil || len(result.Choices) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	amount := countResults(result.Choices)
	rates = append(rates, amount)
	result.Rate = getRate(amount, rates)
	result.Correct = amount

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&result); err != nil {
		panic(err)
	}
}

func getRate(amount int, rates []int) float64 {
	// first result can be the best possible
	minAnswers := 2
	maxRate := 100
	if len(rates) <= minAnswers && amount == NumQuests {
		return maxRate
	}

	var lessThan = 0
	//start at one since 'slice#append' adds to the tail leaving one element always behind
	for i := 1; i < len(rates); i++ {
		if rates[i] < amount {
			lessThan++
		}
	}

	//minus one since 'slice#append' adds to the tail leaving one element always behind
	return math.Round(((float64(lessThan) / float64((len(rates) - 1))) * 100))
}

func countResults(results []string) int {
	var corr = 0
	for idx, result := range results {
		if expectedRes[idx] == result {
			corr++
		}
	}
	return corr
}

func seedQuestions() {
	questions = make([]Question, NumQuests)
	questions[0] = Question{
		Content: "What is the national animal of China?",
		Options: [4]string{"Giant Panda", "Panda", "Red Panda", "Polar Bear"},
	}
	questions[1] = Question{
		Content: "What is the noisiest city in the world?",
		Options: [4]string{"Lisbon", "London", "New York", "Hong Kong"},
	}
	questions[2] = Question{
		Content: "What is a very cold part of Russia?",
		Options: [4]string{"Magadan", "Krasnoyarsk", "Siberia", "Moscow"},
	}
}
