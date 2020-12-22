package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Roll struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var rolls []Roll

func getRoll(w http.ResponseWriter, r *http.Request) {
	//Content-Type headeer indicates if we are working with text/html
	//or text/json
	//This helps the client in processing the response body correctly
	w.Header().Set("Content-Type", "application/json")

	//mux.Vars is setting the params variable from the http response we
	//are passing it
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func getAllRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Here we are creating a new instance of our Roll struct
	var newRoll Roll

	//Getting the data from the HTTP request
	//We are using decoder to read data from the request and put it in
	//"newRoll"
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)

	rolls = append(rolls, newRoll)

	json.NewEncoder(w).Encode(newRoll)
}

func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls = append(rolls, newRoll)
			json.NewEncoder(w).Encode(newRoll)
			return
		}
	}
}

func deleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(rolls)
}

func main() {
	rolls = append(rolls, Roll{ID: "1", Name: "Spicy Tuna Roll"})
	fmt.Println(rolls)
	//initializes router
	router := mux.NewRouter()

	//endpoints
	router.HandleFunc("/sushiRolls", getAllRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	//log.Fatal() throws an error if it fails
	//ListenAndServe() sets up code to run on server
	log.Fatal(http.ListenAndServe(":5000", router))
}
