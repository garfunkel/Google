// Package geocode implements simple geocoding functions.
package geocoding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// GoogleGeocodeURL is the base URL used for the geocoding API.
	GoogleGeocodeURL = "https://maps.googleapis.com/maps/api/geocode/json"
)

// Result is the type specifying geocoding results.
type Result struct {
	AddressComponents []struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	} `json:"address_components"`
	FormattedAddress string `json:"formatted_address"`
	Geometry         struct {
		Location struct {
			Latitude  float64 `json:"lat"`
			Longitude float64 `json:"lng"`
		} `json:"location"`
		LocationType string `json:"location_type"`
		ViewPort     struct {
			NorthEast struct {
				Latitude  float64 `json:"lat"`
				Longitude float64 `json:"lng"`
			} `json:"northeast"`
			SouthWest struct {
				Latitude  float64 `json:"lat"`
				Longitude float64 `json:"lng"`
			} `json:"southwest"`
		} `json:"viewport"`
	} `json:"geometry"`
	Types []string `json:"types"`
}

// Info is the type containing results and success status of geocoding.
type Info struct {
	Results []Result `json:"results"`
	Status  string   `json:"status"`
}

// request generates the request and returns the results.
func request(params string) (info *Info, err error) {
	url := fmt.Sprintf("%s?sensor=false&%s", GoogleGeocodeURL, params)
	response, err := http.Get(url)

	if err != nil {
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	info = new(Info)
	err = json.Unmarshal(data, info)

	return
}

// Geocode performs a forward geocoding request, converting an address into latlng.
func Geocode(address string) (info *Info, err error) {
	params := fmt.Sprintf("address=%v", url.QueryEscape(address))

	return request(params)
}

// ReverseGeocode performs a reverse geocoding request, converting latlng into an address.
func ReverseGeocode(latitude, longitude float64) (info *Info, err error) {
	params := fmt.Sprintf("latlng=%v,%v", latitude, longitude)

	return request(params)
}
