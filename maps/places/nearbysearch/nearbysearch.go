package nearbysearch

import (
	"encoding/json"
	"fmt"
	"github.com/garfunkel/go-google/maps/places"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	APIURL = "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
)

type PipeList []string

type RequiredParams struct {
	APIKey   string          `url:"key"`
	Location places.Location `url:"location"`
}

type OptionalRadiusParam struct {
	Radius int `url:"radius"`
}

type OptionalRankByParam struct {
	RankBy string `url:"rankby"`
}

type OptionalTypesParam struct {
	Types PipeList `url:"types"`
}

type OptionalKeywordParam struct {
	Keyword string `url:"keyword"`
}

type OptionalLanguageParam struct {
	Language string `url:"language"`
}

type OptionalMinPriceParam struct {
	MinPrice int `url:"minprice"`
}

type OptionalMaxPriceParam struct {
	MaxPrice int `url:"maxprice"`
}

type OptionalNameParam struct {
	Name string `url:"name"`
}

type OptionalOpenNowParam struct {
	OpenNow bool `url:"opennow"`
}

type OptionalPageTokenParam struct {
	PageToken string `url:"pagetoken"`
}

type OptionalZagatSelectedParam struct {
	ZagatSelected bool `url:"zagatselected"`
}

func (pipeList PipeList) EncodeValues(key string, v *url.Values) error {
	v.Set(key, strings.Join([]string(pipeList), "|"))

	return nil
}

func NearbySearch(requiredParams *RequiredParams, optionalParams ...interface{}) (searchResponse *places.SearchResponse, err error) {
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

	searchResponse = new(places.SearchResponse)

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
