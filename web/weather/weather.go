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
		Kelvin  float64 `json:"temp"`
		Celsius float64
	} `json:"main"`
}

type api struct {
	key string
}

func NewApi(key string) api {
	return api{key: key}
}

type apiurl string

func (a api) Url() apiurl {
	return apiurl("http://api.openweathermap.org/data/2.5/weather?APPID=" + a.key)
}

type query struct {
	url string
}

func (a apiurl) ByName(cityName string) query {
	return query{url: string(a) + "&q=" + cityName}
}

func (a apiurl) ById(cityId string) query {
	return query{url: string(a) + "&id=" + cityId}
}

// Query returns data from OpenWeatherMap for a given city
func (q query) Query() (Data, error) {
	resp, err := http.Get(q.url)
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
