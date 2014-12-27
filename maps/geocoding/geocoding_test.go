package geocoding

import (
	"log"
	"testing"
)

func TestGeocode(t *testing.T) {
	address := "48 Pirrama Road Pyrmont NSW Australia"
	info, err := Geocode(address)

	if err != nil {
		log.Fatal(err)
	}

	if info.Status != "OK" {
		log.Fatalf("Geocode failed: %v", info.Status)
	}

	info, err = ReverseGeocode(info.Results[0].Geometry.Location.Latitude,
		info.Results[0].Geometry.Location.Longitude)

	if err != nil {
		log.Fatal(err)
	}

	if info.Status != "OK" {
		log.Fatalf("Reverse geocode failed: %v", info.Status)
	}
}
