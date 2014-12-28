// Package distancematrix computes a distance matrix between points using Google's API.
package distancematrix

import (
	"encoding/json"
	"fmt"
	"github.com/garfunkel/go-google/maps"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

const (
	// APIURL is the URL to the API call.
	APIURL = "https://maps.googleapis.com/maps/api/distancematrix/json"
)

// RequiredParams represents params that must be given in a request.
type RequiredParams struct {
	Origins      maps.Locations `url:"origins"`
	Destinations maps.Locations `url:"destinations"`
}

// OptionalKeyParam optional param.
type OptionalKeyParam struct {
	APIKey string `url:"key"`
}

// OptionalModeParam optional param.
type OptionalModeParam struct {
	Mode string `url:"mode"`
}

// OptionalLanguageParam optional param.
type OptionalLanguageParam struct {
	Language string `url:"language"`
}

// OptionalAvoidParam optional param.
type OptionalAvoidParam struct {
	Avoid string `url:"avoid"`
}

// OptionalUnitsParam optional param.
type OptionalUnitsParam struct {
	Units string `url:"units"`
}

// OptionalDepartureTimeParam optional param.
type OptionalDepartureTimeParam struct {
	DepartureTime string `url:"departure_time"`
}

// Response object returned from API call.
type Response struct {
	OriginAddresses      []string `json:"origin_addresses"`
	DestinationAddresses []string `json:"destination_addresses"`
	Status               string   `json:"status"`
	Rows                 []struct {
		Elements []struct {
			Status   string `json:"status"`
			Duration struct {
				Value int    `json:"value"`
				Text  string `json:"text"`
			} `json:"duration"`
			Distance struct {
				Value int    `json:"value"`
				Text  string `json:"text"`
			} `json:"distance"`
		} `json:"elements"`
	} `json:"rows"`
}

// DistanceMatrix computes the distance matrix using Google's API.
func DistanceMatrix(requiredParams *RequiredParams, optionalParams ...interface{}) (matrix *Response, err error) {
	matrix = new(Response)

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

	url := fmt.Sprintf("%v?%v", APIURL, params)
	response, err := http.Get(url)

	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, matrix)

	return
}
