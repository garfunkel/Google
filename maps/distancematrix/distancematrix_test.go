package distancematrix

import (
	"github.com/garfunkel/go-google/maps"
	"log"
	"testing"
)

func TestDistanceMatrix(t *testing.T) {
	requiredParams := RequiredParams{
		Origins: maps.Locations{
			maps.AddressLocation{
				Address: "12 Hunter Street Sydney, NSW, Australia",
			},
			maps.LatLngLocation{
				Latitude:  -33.859235,
				Longitude: 151.068028,
			},
		},
		Destinations: maps.Locations{
			maps.LatLngLocation{
				Latitude:  -33.881059,
				Longitude: 151.174539,
			},
		},
	}

	response, err := DistanceMatrix(&requiredParams)

	if err != nil {
		log.Fatal(err)
	}

	if response.Status != "OK" {
		log.Fatal(response.Status)
	}
}
