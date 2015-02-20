package google

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/drborges/geocoder/providers/google"
    "net/http"
)

func TestReverseGeocodeFromNewGoogleGeocoder(t *testing.T) {
    geocoder := google.NewGoogleGeocoder()

    res, _ := geocoder.ReverseGeocode(47.6064, -122.330803)

    address, _ := google.AddressMapper(res)

    assert.Equal(t, "Seattle", address.City)
    assert.Equal(t, "WA", address.State)
    assert.Equal(t, "US", address.Country)
}

func TestReverseGeocodeFromNewGoogleGeocoderWithHttpProvider(t *testing.T) {
    geocoder := google.NewGoogleGeocoderWithHttpProvider(&http.Client{})

    res, _ := geocoder.ReverseGeocode(47.6064, -122.330803)

    address, _ := google.AddressMapper(res)

    assert.Equal(t, "Seattle", address.City)
    assert.Equal(t, "WA", address.State)
    assert.Equal(t, "US", address.Country)
}

func TestReverseGeocodeFromGoogleCoder(t *testing.T) {
    geocoder := google.GoogleGeocoder{
        HttpClient: &http.Client{},
        ReverseGeocodeEndpoint: "https://maps.googleapis.com/maps/api/geocode/json?latlng=%v,%v",
    }

    res, _ := geocoder.ReverseGeocode(47.6064, -122.330803)

    address, _ := google.AddressMapper(res)

    assert.Equal(t, "Seattle", address.City)
    assert.Equal(t, "WA", address.State)
    assert.Equal(t, "US", address.Country)
}
