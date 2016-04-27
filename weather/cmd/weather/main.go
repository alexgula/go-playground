package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/alexgula/go-playground/weather"
)

func getWeather(key string, city string) {
	data, err := weather.NewApi(key, log.New(ioutil.Discard, "LOG: ", log.Ldate|log.Ltime)).Url().ByName(city).Query()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(data)
	fmt.Printf("%s, %s %.1f\u00B0C (%.1f\u00B0C - %.1f\u00B0C) %.0f%% %.0fhPa\n",
		data.Name,
		data.Sys.Country,
		data.Main.Temp,
		data.Main.TempMin,
		data.Main.TempMax,
		data.Main.Humidity,
		data.Main.Pressure)
}

func runWeather(interval time.Duration, key string, city string) sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			getWeather(key, city)
			time.Sleep(interval)
		}
		wg.Done()
	}()
	return wg
}

func main() {
	fmt.Println("Starting...")

	key := flag.String("key", "", "openweathermap.org API key")
	city := flag.String("city", "", "openweathermap.org city name")
	intervalFlag := flag.String("interval", "10m", "refresh interval")
	flag.Parse()

	interval, err := time.ParseDuration(*intervalFlag)
	if err != nil {
		log.Fatalf("refresh interval is not in correct format: %s", *intervalFlag)
	}

	wg := runWeather(interval, *key, *city)
	wg.Wait()
}
