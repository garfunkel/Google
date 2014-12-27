package nearbysearch

import (
	"errors"
	"github.com/garfunkel/go-google/maps/places"
	"log"
	"testing"
)

func TestGetClosestPlaces(t *testing.T) {
	requiredParams := RequiredParams{
		APIKey: "AIzaSyC50lfM-BNpgJMXesZ9qV4Jx6ubTMmwwxA",
		Location: places.Location{
			Latitude:  -33.859235,
			Longitude: 151.068028,
		},
	}

	rankByParam := OptionalRankByParam{
		RankBy: "distance",
	}

	typesParam := OptionalTypesParam{
		Types: PipeList{"train_station", "bus_station"},
	}

	keywordParam := OptionalKeywordParam{
		Keyword: "flemington",
	}

	searchResponse, err := NearbySearch(&requiredParams, &rankByParam, &typesParam, &keywordParam)

	if err != nil {
		log.Fatal(err)
	}

	if searchResponse.Status != "OK" {
		log.Fatal(errors.New(searchResponse.Status))
	}
}
