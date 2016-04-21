package main

import (
	"fmt"

	"github.com/alexgula/go-playground/weather"
)

func getWeather(key string, city string) (weather.Data, error) {
	data, err := weather.NewApi(key).Url().ByName(city).Query()
	if err != nil {
		return weather.Data{}, err
	}

	data.City = city
	data.Main.Celsius = data.Main.Kelvin - 273.15

	return data, nil
}

func main() {
	data, err := getWeather("602e5a1c5cb62e61550a72adf8726063", "Gdansk")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%s %v\u00B0C\n", data.City, data.Main.Celsius)
}
