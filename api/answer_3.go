package main

import (
	"net/http"
	"net/url"
	"strings"

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

	// Format HTTP Request to data provider
	params := url.Values{}
	params.Add("orderByFields", "DATE")
	params.Add("outFields", "*")
	params.Add("where", "STATE = '"+state+"' AND Year = "+vars["year"])
	params.Add("f", "json")

	// Process the Request to and Response from data provider
	var raw ArcgisNOAASEDResponse
	JSONRequest(ArcgisNOAASEDAPIURL+"?"+params.Encode(), &raw)

	// Transform data
	var data []ArcgisNOAASEDResponseFeatureAttributeTransformed
	for _, a := range raw.Features {
		data = append(data, ArcgisNOAASEDResponseFeatureAttributeTransformed{
			ID:               a.Attributes.ID,
			County:           a.Attributes.County,
			Damage:           parseArcgisNOAASEDDamage(a.Attributes.Damage),
			Date:             parseArcgisNOAASEDDate(a.Attributes.Date),
			DeathsDirect:     a.Attributes.DeathsDirect,
			DeathsIndirect:   a.Attributes.DeathsIndirect,
			Narrative:        a.Attributes.Narrative,
			Event:            a.Attributes.Type,
			InjuriesDirect:   a.Attributes.InjuriesDirect,
			InjuriesIndirect: a.Attributes.InjuriesIndirect,
			State:            a.Attributes.State,
			Year:             a.Attributes.Year,
		})
	}

	// Output as JSON
	outputJSON(w, data)
}
