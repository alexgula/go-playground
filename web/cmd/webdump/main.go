package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Request</h1>")
	fmt.Fprintf(w, "<pre>"+html.EscapeString(spew.Sdump(r))+"</pre>")
	fmt.Fprintf(w, "<h1>Response</h1>")
	fmt.Fprintf(w, "<pre>"+html.EscapeString(spew.Sdump(w))+"</pre>")
}

func main() {
	port := 10002
	fmt.Printf("Started server on %d\n", port)
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.FileServer(FS(false)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
