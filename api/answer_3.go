package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func apiAnswer3(w http.ResponseWriter, r *http.Request) {
	// Allow CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    if r.Method == http.MethodOptions {
        return
    }
	// Parse the HTTP Request to this service
	vars := mux.Vars(r)
	state := strings.ToUpper(vars["state"]) // State has to be uppercase
	year, _ := strconv.Atoi(vars["year"])
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(1, 0, 0)
	// Dates need to be in "YYYY-MM-DD hh:mm:ss" format
	// This is not an ISO or RFC standard, so manual formatting is required
	fmtStart := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", start.Year(), start.Month(), start.Day(), start.Hour(), start.Minute(), start.Second())
	fmtEnd := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", end.Year(), end.Month(), end.Day(), end.Hour(), end.Minute(), end.Second())

	// Format HTTP Request to data provider
	params := url.Values{}
	params.Add("outFields", "*")
	params.Add("where", "STATE = '"+state+"' AND DATE >= TIMESTAMP '"+fmtStart+"' AND DATE <= TIMESTAMP '"+fmtEnd+"'")
	params.Add("f", "json")

	// Process the Request to and Response from data provider
	var data ArcgisNOAASEDResponse
	JSONRequest(ArcgisNOAASEDAPIURL+"?"+params.Encode(), &data)

	// Output as JSON
	outputJSON(w, data.Features)
}
