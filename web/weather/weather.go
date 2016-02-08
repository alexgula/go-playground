package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(fmt.Sprintf("Could not close the response body: %#v", err))
		}
	}()

	var d Data

	log.Printf("%#v", resp)

	if err := json.NewDecoder(io.TeeReader(resp.Body, &logWriter{})).Decode(&d); err != nil {
		return Data{}, err
	}

	fmt.Fprintln(os.Stderr)

	return d, nil
}

type logWriter struct {
}

func (l *logWriter) Write(p []byte) (n int, err error) {
	log.Printf("%s", p)
	return
}
