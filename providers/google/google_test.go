package google

import (
	"github.com/drborges/geocoder/providers/google"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestReverseGeocodeFromNewGoogleGeocoder(t *testing.T) {
	geocoder := google.NewGeocoder()

	res, _ := geocoder.ReverseGeocode(47.6064, -122.330803)

	address, _ := google.AddressMapper(res)

	assert.Equal(t, "Seattle", address.City)
	assert.Equal(t, "WA", address.State)
	assert.Equal(t, "US", address.Country)
}

func TestReverseGeocodeFromNewGoogleGeocoderWithHttpProvider(t *testing.T) {
	geocoder := google.NewGeocoderWithHttpProvider(&http.Client{})

	res, _ := geocoder.ReverseGeocode(47.6064, -122.330803)

	address, _ := google.AddressMapper(res)

	assert.Equal(t, "Seattle", address.City)
	assert.Equal(t, "WA", address.State)
	assert.Equal(t, "US", address.Country)
}

func TestReverseGeocodeFromGoogleCoder(t *testing.T) {
	geocoder := google.Geocoder{
		HttpClient:             &http.Client{},
		ReverseGeocodeEndpoint: "https://maps.googleapis.com/maps/api/geocode/json?latlng=%v,%v",
	}

	res, _ := geocoder.ReverseGeocode(47.6064, -122.330803)

	address, _ := google.AddressMapper(res)

	assert.Equal(t, "Seattle", address.City)
	assert.Equal(t, "WA", address.State)
	assert.Equal(t, "US", address.Country)
}
