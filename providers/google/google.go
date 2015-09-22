package google

import (
	"encoding/json"
	"fmt"
	"github.com/drborges/geocoder"
	"io/ioutil"
	"net/http"
	"reflect"
)

const (
	ReverseGeocodeEndpoint = "https://maps.googleapis.com/maps/api/geocode/json?latlng=%v,%v"
)

type Geocoder struct {
	HttpClient             *http.Client
	ReverseGeocodeEndpoint string
}

func NewGeocoder() geocoder.Geocoder {
	return &Geocoder{
		HttpClient:             &http.Client{},
		ReverseGeocodeEndpoint: ReverseGeocodeEndpoint,
	}
}

func NewGeocoderWithHttpProvider(c *http.Client) geocoder.Geocoder {
	return &Geocoder{
		HttpClient:             c,
		ReverseGeocodeEndpoint: ReverseGeocodeEndpoint,
	}
}

func (geo *Geocoder) ReverseGeocode(lat float64, lng float64) (*http.Response, error) {
	return geo.HttpClient.Get(fmt.Sprintf(geo.ReverseGeocodeEndpoint, lat, lng))
}

type Address struct {
	Country string `field:"short_name" type:"country"`
	State   string `field:"short_name" type:"administrative_area_level_1"`
	City    string `field:"short_name" type:"locality"`

	FullCountry string `field:"long_name" type:"country"`
	FullState   string `field:"long_name" type:"administrative_area_level_1"`
	FullCity    string `field:"long_name" type:"locality"`
}

func ReadResponse(res *http.Response, dst interface{}) error {
	jsonData := readJson(res)

	elem := reflect.TypeOf(dst).Elem()
	elemValue := reflect.ValueOf(dst).Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fieldValue := elemValue.Field(i)
		mapField(jsonData, fieldValue, field.Tag.Get("field"), field.Tag.Get("type"))
	}

	return nil
}

func readJson(res *http.Response) map[string]interface{} {
	jsonData := map[string]interface{}{}

	if body, err := ioutil.ReadAll(res.Body); err == nil {
		json.Unmarshal(body, &jsonData)
	}

	return jsonData
}

// TODO refactor this mess
// Type assertions makes readability even harder.
// Perhaps it would be better to map all fields at once rather than one by one.
func mapField(jsonData map[string]interface{}, field reflect.Value, gmapsFieldName, gmapsFieldType string) {
	for _, result := range jsonData["results"].([]interface{}) {
		for _, component := range result.(map[string]interface{})["address_components"].([]interface{}) {
			for _, componentType := range component.(map[string]interface{})["types"].([]interface{}) {
				if componentType.(string) == gmapsFieldType {
					field.SetString(component.(map[string]interface{})[gmapsFieldName].(string))
					return
				}
			}
		}
	}
}
