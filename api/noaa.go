package main

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
