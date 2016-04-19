package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/alexgula/go-playground/weather"
)

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 2)[1]

	data, err := weather.NewApi("602e5a1c5cb62e61550a72adf8726063").Url().ByName(city).Query()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.City = city

	data.Main.Celsius = data.Main.Kelvin - 273.15

	w.Header().Set("Content-Type", "application/json, charset=utf-8")

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	port := 10000
	fmt.Printf("Started server on %d\n", port)
	http.HandleFunc("/", weatherHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
