// Package maps implements common Google Maps API features.
package maps

import (
	"fmt"
	"net/url"
	"strings"
)

// Geometry represents location geometry.
type Geometry struct {
	Location LatLngLocation `json:"location"`
}

// Location interface for various ways to structure locations.
type Location interface {
	EncodeValues(key string, v *url.Values) error
}

// Locations is a slice of Location interfaces.
type Locations []Location

// AddressLocation represents a location as an address string.
type AddressLocation struct {
	Address string `json:"address"`
}

// LatLngLocation represents a location as latlng coordinates.
type LatLngLocation struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// EncodeValues encodes AddressLocation into URL form.
func (location AddressLocation) EncodeValues(key string, v *url.Values) error {
	values, _ := (*v)[key]

	values = append(values, location.Address)

	v.Set(key, strings.Join(values, "|"))

	return nil
}

// EncodeValues encodes LatLngLocation into URL form.
func (location LatLngLocation) EncodeValues(key string, v *url.Values) error {
	values, _ := (*v)[key]

	values = append(values, fmt.Sprintf("%v,%v", location.Latitude, location.Longitude))

	v.Set(key, strings.Join(values, "|"))

	return nil
}

// EncodeValues encodes a slice of Location interfaces into URL form.
func (locations Locations) EncodeValues(key string, v *url.Values) error {
	for _, location := range locations {
		if err := location.EncodeValues(key, v); err != nil {
			return err
		}
	}

	return nil
}
