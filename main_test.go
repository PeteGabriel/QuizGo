package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const expectedQuiz = `[{"content":"What is the national animal of China?","options":["Giant Panda","Panda","Red Panda","Polar Bear"]},{"content":"What is the noisiest city in the world?","options":["Lisbon","London","New York","Hong Kong"]},{"content":"What is a very cold part of Russia?","options":["Magadan","Krasnoyarsk","Siberia","Moscow"]}]`

func TestGetQuiz_ExpectSetOfQuestions(t *testing.T) {
	//setup
	req, err := http.NewRequest("GET", "/quiz", nil)
	if err != nil {
		t.Fatal(err)
	}

	//act
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetQuiz)
	handler.ServeHTTP(rr, req)

	//assert
	checkResponseCode(t, http.StatusOK, rr.Code)
	checkMediatype(t, rr.Header().Get("Content-Type"))

	if strings.TrimRight(rr.Body.String(), "\n") != expectedQuiz {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedQuiz)
	}
}

func TestSaveQuiz_ExpectValidResultAndTopRate(t *testing.T) {
	//setup
	expectedRes := `{"choices":["Giant Panda","Hong Kong","Siberia"],"successRate":100,"correct":3}`
	body := `{"choices":["Giant Panda", "Hong Kong", "Siberia"]}`
	req, err := http.NewRequest("POST", "/quiz/results", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	//act
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveResults)
	handler.ServeHTTP(rr, req)

	//assert
	checkResponseCode(t, http.StatusCreated, rr.Code)
	checkMediatype(t, rr.Header().Get("Content-Type"))

	if strings.TrimRight(rr.Body.String(), "\n") != expectedRes {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedRes)
	}
}

func TestSaveQuizTwice_ExpectValidResultAnd50Rate(t *testing.T) {
	//setup
	expectedRes := `{"choices":["Giant Panda","Hong Kong","Siberia"],"successRate":50}`

	body := `{"choices":["Panda", "Hong Kong", "Siberia"]}`
	req, err := http.NewRequest("POST", "/quiz/results", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	//second request
	body2 := `{"choices":["Giant Panda", "Hong Kong", "Siberia"]}`
	req2, err := http.NewRequest("POST", "/quiz/results", strings.NewReader(body2))
	if err != nil {
		t.Fatal(err)
	}

	//act
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveResults)
	handler.ServeHTTP(rr, req)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req2)

	//assert
	if strings.TrimRight(rr.Body.String(), "\n") != expectedRes {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedRes)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkMediatype(t *testing.T, actual string) {
	if actual != "application/json" {
		t.Errorf("handler returned wrong mediatype: got %v want %v",
			actual, "application/json")
	}
}
