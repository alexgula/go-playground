package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Hello, %s!</p>", r.URL.Path[1:])
	fmt.Printf("Hello, %s!\n", r.URL.Path[1:])
	fmt.Fprintf(w, "<p>I'm running on %s with an %s CPU</p>", runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
