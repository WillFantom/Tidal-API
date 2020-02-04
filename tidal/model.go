package tidal

type Tidal struct {
	SessionID   string `json:"sessionId"`
	CountryCode string `json:"countryCode"`
	UserID      uint   `json:"userId"`
	quality     string
}

type Artist struct {
}
