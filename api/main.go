package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// Initialize HTTP router
	router := mux.NewRouter()
	api := router.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/answer/1/{slug}/{year}", apiAnswer1)
	api.HandleFunc("/answer/3/{state}/{year}", apiAnswer3)

	// Configure HTTP Server
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// Run
	log.Fatal(srv.ListenAndServe())
}
