package google

import (
    "testing"
    "net/http"
    "github.com/stretchr/testify/assert"
    "github.com/drborges/geocoder/google"
)

func TestReverseGeocode(t *testing.T) {
    geocoder := google.NewGeoCoderWithHttpClient(http.Client{})

    res, _ := geocoder.ReverseGeocode(47.6064, -122.330803)

    address, _ := google.GoogleAddressMapper(res)

    assert.Equal(t, "Seattle", address.City)
    assert.Equal(t, "WA", address.State)
    assert.Equal(t, "US", address.Country)
}