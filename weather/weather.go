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
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Dt   int64 `json:"dt"`
	Main struct {
		Humidity float64 `json:"humidity"` // %
		Pressure float64 `json:"pressure"` // hPa
		Temp     float64 `json:"temp"`     // Celsius
		TempMin  float64 `json:"temp_min"` // Celsius
		TempMax  float64 `json:"temp_max"` // Celsius
	} `json:"main"`
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	} `json:"sys"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed  float64 `json:"speed"` // m/s
		Degree float64 `json:"deg"`   // Degree
	} `json:"wind"`
}

type api struct {
	key    string
	logger *log.Logger
}

func NewApi(key string, logger *log.Logger) api {
	return api{key: key, logger: logger}
}

type apiurl struct {
	api
	url string
}

func (a api) Url() apiurl {
	return apiurl{api: a, url: "http://api.openweathermap.org/data/2.5/weather?units=metric&APPID=" + a.key}
}

type query struct {
	apiurl
}

func (a apiurl) ByName(cityName string) query {
	a.url = a.url + "&q=" + cityName
	return query{apiurl: a}
}

func (a apiurl) ById(cityId string) query {
	a.url = a.url + "&id=" + cityId
	return query{apiurl: a}
}

type queryLogWriter struct {
	query
}

func (l *queryLogWriter) Write(p []byte) (n int, err error) {
	l.logger.Printf("%s", p)
	return
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
			q.logger.Fatal(fmt.Sprintf("Could not close the response body: %#v", err))
		}
	}()

	var d Data

	q.logger.Printf("%#v", resp)

	if err := json.NewDecoder(io.TeeReader(resp.Body, &queryLogWriter{query: q})).Decode(&d); err != nil {
		return Data{}, err
	}

	fmt.Fprintln(os.Stderr)

	return d, nil
}
