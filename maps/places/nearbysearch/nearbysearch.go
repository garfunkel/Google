// Package nearbysearch implements Google Place's Nearby Search API.
package nearbysearch

import (
	"encoding/json"
	"fmt"
	"github.com/garfunkel/go-google/maps"
	"github.com/garfunkel/go-google/maps/places"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// APIURL is the URL to the API call.
	APIURL = "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
)

// PipeList represents a slice of strings to be separated by a pipe symbol.
type PipeList []string

// RequiredParams represents params required by the API.
type RequiredParams struct {
	APIKey   string        `url:"key"`
	Location maps.Location `url:"location"`
}

// OptionalRadiusParam optional param.
type OptionalRadiusParam struct {
	Radius int `url:"radius"`
}

// OptionalRankByParam optional param.
type OptionalRankByParam struct {
	RankBy string `url:"rankby"`
}

// OptionalTypesParam optional param.
type OptionalTypesParam struct {
	Types PipeList `url:"types"`
}

// OptionalKeywordParam optional param.
type OptionalKeywordParam struct {
	Keyword string `url:"keyword"`
}

// OptionalLanguageParam optional param.
type OptionalLanguageParam struct {
	Language string `url:"language"`
}

// OptionalMinPriceParam optional param.
type OptionalMinPriceParam struct {
	MinPrice int `url:"minprice"`
}

// OptionalMaxPriceParam optional param.
type OptionalMaxPriceParam struct {
	MaxPrice int `url:"maxprice"`
}

// OptionalNameParam optional param.
type OptionalNameParam struct {
	Name string `url:"name"`
}

// OptionalOpenNowParam optional param.
type OptionalOpenNowParam struct {
	OpenNow bool `url:"opennow"`
}

// OptionalPageTokenParam optional param.
type OptionalPageTokenParam struct {
	PageToken string `url:"pagetoken"`
}

// OptionalZagatSelectedParam optional param.
type OptionalZagatSelectedParam struct {
	ZagatSelected bool `url:"zagatselected"`
}

// EncodeValues encodes PipeList into into URL form.
func (pipeList PipeList) EncodeValues(key string, v *url.Values) error {
	v.Set(key, strings.Join([]string(pipeList), "|"))

	return nil
}

// NearbySearch performs a search using Google's API.
func NearbySearch(requiredParams *RequiredParams, optionalParams ...interface{}) (searchResponse *places.Response, err error) {
	values, err := query.Values(requiredParams)

	if err != nil {
		return
	}

	params := values.Encode()

	for _, optionalParam := range optionalParams {
		values, err = query.Values(optionalParam)

		if err != nil {
			return
		}

		params += fmt.Sprintf("&%v", values.Encode())
	}

	searchResponse = new(places.Response)

	url := fmt.Sprintf("%v?%v", APIURL, params)
	response, err := http.Get(url)

	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, searchResponse)

	return
}
