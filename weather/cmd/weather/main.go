package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/alexgula/go-playground/weather"
)

type weatherClient struct {
	logger *log.Logger
	key    string
	city   string
}

func (wc weatherClient) get() {
	data, err := weather.NewApi(wc.key, wc.logger).Url().ByName(wc.city).Query()
	if err != nil {
		log.Fatal(err)
		wc.logger.Fatal(err)
	}

	date := time.Unix(data.Dt, 0)

	wc.logger.Print(data)
	fmt.Printf("%v: %s, %s %.1f\u00B0C (%.1f\u00B0C - %.1f\u00B0C) %.0f%% %.0fhPa\n",
		date,
		data.Name,
		data.Sys.Country,
		data.Main.Temp,
		data.Main.TempMin,
		data.Main.TempMax,
		data.Main.Humidity,
		data.Main.Pressure)
}

func (wc weatherClient) poll(interval time.Duration) sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			wc.get()
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
	loggerName := flag.String("log", "", "logger")
	flag.Parse()

	interval, err := time.ParseDuration(*intervalFlag)
	if err != nil {
		log.Fatalf("refresh interval is not in correct format: %s", *intervalFlag)
	}

	logWriter := ioutil.Discard
	if *loggerName == "stdout" {
		logWriter = os.Stdout
	}
	logger := log.New(logWriter, "LOG: ", log.Ldate|log.Ltime)
	wc := weatherClient{logger: logger, key: *key, city: *city}
	wg := wc.poll(interval)
	wg.Wait()
}
