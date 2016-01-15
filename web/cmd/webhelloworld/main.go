package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	port := 10001
	fmt.Printf("Starting server on %d\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, I don't know you, but you are asking for %q, right?", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
