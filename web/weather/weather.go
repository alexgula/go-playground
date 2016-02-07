package weather

import (
	"encoding/json"
	"net/http"
)

// Data contains weather results from OpenWeatherMap along with additional info
type Data struct {
	Name string `json:"name"`
	City string `json:"city"`
	Main struct {
		Kelvin  float64 `json:"kelvin"`
		Celsius float64 `json:"celsius"`
	} `json:"main"`
}

// Query returns data from OpenWeatherMap for a given city
func Query(city string) (Data, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city)
	if err != nil {
		return Data{}, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var d Data

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return Data{}, err
	}

	return d, nil
}
