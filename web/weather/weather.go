package weather

import (
	"encoding/json"
	"net/http"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin  float64 `json:"temp"`
		Celsius float64
	} `json:"main"`
}

func Query(city string) (WeatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city)
	if err != nil {
		return WeatherData{}, err
	}

	defer resp.Body.Close()

	var d WeatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return WeatherData{}, err
	}

	return d, nil
}
