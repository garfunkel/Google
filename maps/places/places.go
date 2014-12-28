// Package places implements Google's Places API.
package places

import (
	"github.com/garfunkel/go-google/maps"
)

// Result structure for search results.
type Result struct {
	Geometry  maps.Geometry `json:"geometry"`
	Icon      string        `json:"icon"`
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	PlaceID   string        `json:"place_id"`
	Reference string        `json:"reference"`
	Scrope    string        `json:"scope"`
	Types     []string      `json:"types"`
	Vicinity  string        `json:"vicinity"`
	Photos    []struct {
		Height           int
		Width            int
		HTMLAttributions []string `json:"html_attributions"`
		Reference        string   `json:"photo_reference"`
	} `json:"photos"`
	OpeningHours struct {
		OpenNow     bool     `json:"open_now"`
		WeekdayText []string `json:"weekday_text"`
	} `json:"opening_hours"`
	Rating         float64 `json:"rating"`
	PriceLevel     int     `json:"price_level"`
	AlternativeIDs []struct {
		PlaceID string `json:"place_id"`
		Scope   string `json:"scope"`
	} `json:"alt_ids"`
}

// Response structure for a group of search results.
type Response struct {
	HTMLAttributions []string `json:"html_attributions"`
	NextPageToken    string   `json:"next_page_token"`
	Results          []Result `json:"results"`
	Status           string   `json:"status"`
}
