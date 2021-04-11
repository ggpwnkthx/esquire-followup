package main

import (
	"strconv"
	"strings"
	"time"
)

// ArcGIS / NOAA Constants
const ArcgisNOAASEDAPIURL = "https://services.arcgis.com/jIL9msH9OI208GCb/arcgis/rest/services/NOAA_Storm_Events_Database_view/FeatureServer/0/query"

// ArcGIS / NOAA Data Models
type ArcgisNOAASEDResponse struct {
	Features []ArcgisNOAASEDResponseFeature `json:"features"`
}
type ArcgisNOAASEDResponseFeature struct {
	Attributes ArcgisNOAASEDResponseFeatureAttribute `json:"attributes"`
}
type ArcgisNOAASEDResponseFeatureAttribute struct {
	County           string `json:"CZ_NAME"`
	Damage           string `json:"DAMAGE_PROPERTY"`
	Date             int    `json:"DATE"`
	DeathsDirect     int    `json:"DEATHS_DIRECT"`
	DeathsIndirect   int    `json:"DEATHS_INDIRECT"`
	Narrative        string `json:"EPISODE_NARRATIVE"`
	Type             string `json:"EVENT_TYPE"`
	InjuriesDirect   int    `json:"INJURIES_DIRECT"`
	InjuriesIndirect int    `json:"INJURIES_INDIRECT"`
	ID               int    `json:"OBJECTID"`
	State            string `json:"STATE"`
	TorFScale        string `json:"TOR_F_SCALE"`
	Year             int    `json:"Year"`
}
type ArcgisNOAASEDResponseFeatureAttributeTransformed struct {
	County           string    `json:"County"`
	Damage           int       `json:"Damage"`
	Date             time.Time `json:"Date"`
	DeathsDirect     int       `json:"DeathsDirect"`
	DeathsIndirect   int       `json:"DeathsIndirect"`
	Narrative        string    `json:"Narative"`
	Event            string    `json:"Event"`
	InjuriesDirect   int       `json:"InjuriesDirect"`
	InjuriesIndirect int       `json:"InjuriesIndirect"`
	ID               int       `json:"ID"`
	State            string    `json:"State"`
	Year             int       `json:"Year"`
}

func parseArcgisNOAASEDDamage(input string) int {
	if input == "" {
		return 0
	}
	number, _ := strconv.ParseFloat(strings.Split(input, "K")[0], 64)
	return int(number * 1000)
}
func parseArcgisNOAASEDDate(input int) time.Time {
	epoch := time.Date(1970, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	return epoch.Add(time.Millisecond * time.Duration(input))
}
