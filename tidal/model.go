package tidal

type Tidal struct {
	SessionID   string `json:"sessionId"`
	CountryCode string `json:"countryCode"`
	UserID      uint   `json:"userId"`
	quality     string
}

type Artist struct {
	ID          uint     `json:"id" mapstructure:"id"`
	Name        string   `json:"name" mapstructure:"name"`
	ArtistTypes []string `json:"artistTypes" mapstructure:"artistTypes"`
	Popularity  int      `json:"popularity" mapstructure:"popularity"`
	ImgURL      string
}

type Album struct {
	ID        uint     `json:"id" mapstructure:"id"`
	Name      string   `json:"title" mapstructure:"title"`
	Artist    Artist   `json:"artist" mapstructure:"artist"`
	Artists   []Artist `json:"artists" mapstructure:"artists"`
	Tracks    int      `json:"numberOfTracks" mapstructure:"numberOfTracks"`
	Duration  int      `json:"duration" mapstructure:"duration"`
	Copyright string   `json:"copyright" mapstructure:"copyright"`
	Explicit  bool     `json:"explicit" mapstructure:"explicit"`
	ImgURL    string
}

type AlbumShort struct {
	ID   uint   `json:"id" mapstructure:"id"`
	Name string `json:"title" mapstructure:"title"`
}

type Track struct {
	ID         uint       `json:"id" mapstructure:"id"`
	Name       string     `json:"title" mapstructure:"title"`
	Duration   int        `json:"duration" mapstructure:"duration"`
	Track      int        `json:"trackNumber" mapstructure:"trackNumber"`
	Disc       int        `json:"volumeNumber" mapstructure:"volumeNumber"`
	Version    int        `json:"version" mapstructure:"version"`
	Popularity int        `json:"popularity" mapstructure:"popularity"`
	Artist     Artist     `json:"artist" mapstructure:"artist"`
	Artists    []Artist   `json:"artists" mapstructure:"artists"`
	Album      AlbumShort `json:"album" mapstructure:"album"`
	Available  bool       `json:"streamReady" mapstructure:"streamReady"`
	Format     string     `json:"type" mapstructure:"type"`
	Explicit   bool       `json:"explicit" mapstructure:"explicit"`
}

type SearchResponse struct {
	Limit      int           `json:"limit"`
	Offset     int           `json:"offset"`
	TotalItems int           `json:"totalNumberOfItems"`
	Items      []interface{} `json:"items"`
}
