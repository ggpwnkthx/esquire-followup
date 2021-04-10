package main

import (
	"time"
)

// Nomics Constants
const NomicsAPIURL = "https://api.nomics.com/v1/"
const NomicsAPIKey = "6b42d899b77efc3270143238a6dd703c"

// Nomics Data Models
type ExchangeRateHistory struct {
	Timestamp time.Time `json:"timestamp"`
	Rate      string    `json:"rate"`
}
