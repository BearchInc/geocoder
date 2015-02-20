package google

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/drborges/geocoder/providers/google"
)

func TestReverseGeocode(t *testing.T) {
    geocoder := google.NewGoogleGeocoder()

    res, _ := geocoder.ReverseGeocode(47.6064, -122.330803)

    address, _ := google.AddressMapper(res)

    assert.Equal(t, "Seattle", address.City)
    assert.Equal(t, "WA", address.State)
    assert.Equal(t, "US", address.Country)
}