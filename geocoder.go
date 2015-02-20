package geocoder

import "net/http"

type GeoLocation struct {
    Lat float64
    Lng float64
}

type Geocoder interface {
    ReverseGeocode(float64, float64) (*http.Response, error)
//    Geocode(string) (GeoLocation, error)
}

type Address struct {
    Country string
    State   string
    City    string
}

type AddressMapper func(*http.Response) Address

var EmptyAddress = Address{}