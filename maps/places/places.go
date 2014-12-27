package places

import (
	"fmt"
	"net/url"
)

type Geometry struct {
	Location Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Result struct {
	Geometry  Geometry `json:"geometry"`
	Icon      string   `json:"icon"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	PlaceID   string   `json:"place_id"`
	Reference string   `json:"reference"`
	Scrope    string   `json:"scope"`
	Types     []string `json:"types"`
	Vicinity  string   `json:"vicinity"`
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

type SearchResponse struct {
	HTMLAttributions []string `json:"html_attributions"`
	NextPageToken    string   `json:"next_page_token"`
	Results          []Result `json:"results"`
	Status           string   `json:"status"`
}

func (location Location) EncodeValues(key string, v *url.Values) error {
	v.Set(key, fmt.Sprintf("%v,%v", location.Latitude, location.Longitude))

	return nil
}
