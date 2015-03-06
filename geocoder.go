package geocoder

import "net/http"

type Geocoder interface {
	ReverseGeocode(float64, float64) (*http.Response, error)
}
