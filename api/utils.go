package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func JSONRequest(url string, data interface{}) {
	// Initilaize HTTP Client
	client := http.Client{
		Timeout: time.Second * 15, // Timeout after 15 seconds
	}
	// Configure Request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}
	// Send Request
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	// Decode Response
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}
}
func outputJSON(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		log.Println(err)
	}
}
