package main

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func apiAnswer1(w http.ResponseWriter, r *http.Request) {
	// Parse the HTTP Request to this service
	vars := mux.Vars(r)
	slug := strings.ToUpper(vars["slug"]) // Slug has to be uppercase
	year, _ := strconv.Atoi(vars["year"])
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(1, 0, 0)

	// Format HTTP Request to data provider
	params := url.Values{}
	params.Add("key", NomicsAPIKey)
	params.Add("currency", slug)
	params.Add("start", start.Format(time.RFC3339))
	params.Add("end", end.Format(time.RFC3339))

	// Process the Request to and Response from data provider
	var data []ExchangeRateHistory
	JSONRequest(NomicsAPIURL+"exchange-rates/history?"+params.Encode(), &data)

	// Output as JSON
	outputJSON(w, data)
}
