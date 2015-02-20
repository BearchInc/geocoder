package google

import (
    "net/http"
    "github.com/drborges/geocoder"
    "encoding/json"
    "io/ioutil"
    "fmt"
)

const ReverseGeocodeEndpoint = "https://maps.googleapis.com/maps/api/geocode/json?latlng=%v,%v"

type GoogleResponse struct {
    Results []GoogleResults `json:"results"`
}

func (res *GoogleResponse) Address() geocoder.Address {
    address := geocoder.Address{}
    for _, result := range res.Results {
        if result.Types[0] == "locality" {
            for _, addrComponent := range result.AddressComponents {
                if addrComponent.Types[0] == "locality" {
                    address.City = addrComponent.ShortName
                } else if addrComponent.Types[0] == "administrative_area_level_1" {
                    address.State = addrComponent.ShortName
                } else if addrComponent.Types[0] == "country" {
                    address.Country = addrComponent.ShortName
                }
            }
        }
    }
    return address
}

type GoogleResults struct {
    Types []string `json:"types"`
    AddressComponents []AddressComponents `json:"address_components"`
}

type AddressComponents struct {
    LongName string  `json:"long_name"`
    ShortName string `json:"short_name"`
    Types []string   `json:"types"`
}

type GoogleGeoCoder struct {
    Http http.Client
}

func NewGeoCoderWithHttpClient(c http.Client) geocoder.Geocoder {
    return &GoogleGeoCoder{c}
}

func (geo *GoogleGeoCoder) ReverseGeocode(lat float64, lng float64) (*http.Response, error) {
    return geo.Http.Get(fmt.Sprintf(ReverseGeocodeEndpoint, lat, lng))
}

func GoogleAddressMapper(res *http.Response) (geocoder.Address, error) {
    if body, err := ioutil.ReadAll(res.Body); err == nil {
        gres := new(GoogleResponse)
        json.Unmarshal(body, &gres)
        return gres.Address(), nil
    } else {
        return geocoder.Address{}, err
    }
}