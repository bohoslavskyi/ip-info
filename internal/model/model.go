package model

type IPInfo struct {
	Status      string  `json:"status"`
	Message     string  `json:"message"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
}

type IPDetails struct {
	IP          string           `json:"ip"`
	Country     string           `json:"country"`
	City        string           `json:"city"`
	Latitude    float64          `json:"latitude"`
	Longitude   float64          `json:"longitude"`
	CurrentTime string           `json:"currentTime"`
	Currencies  []map[string]any `json:"currencies"`
	Err         error            `json:"-"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
