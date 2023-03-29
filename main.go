package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/calculate", calculateFlightPath)
	http.ListenAndServe(":8080", nil)
}

func calculateFlightPath(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into a [][]string slice
	var flights [][]string
	if err := json.NewDecoder(r.Body).Decode(&flights); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Create a map to store the source and destination airports
	airports := make(map[string]string)
	// To find the starting airport as quickly as possible (O(n)), we will store a reverse map of destination -> starting to see which airport is not a destination
	reverseAirports := make(map[string]string)

	// Store the starting and destination airports into map
	for _, flight := range flights {
		airports[flight[0]] = flight[1]
		reverseAirports[flight[1]] = flight[0]
	}

	// Find the starting airport by seeing which key of airports is not present in reverseAirports
	startingAirport := ""
	for airport := range airports {
		if _, ok := reverseAirports[airport]; !ok {
			startingAirport = airport
			break
		}
	}

	// Find the ending airport by iterating over the airports in order until there we reached the final destination
	var endingAirport string
	currentAirport := startingAirport
	for {
		nextAirport, ok := airports[currentAirport]
		if !ok {
			endingAirport = currentAirport
			break
		}
		currentAirport = nextAirport
	}

	// The task instructed to return the starting and ending airports, as opposed to the total flight path
	startingAndEnding := []string{startingAirport, endingAirport}

	json.NewEncoder(w).Encode(startingAndEnding)
}
