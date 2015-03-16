package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
)

type word string

type command struct {
	verb word
}

type greeting struct {
	verb word
	noun word
}

type runtime struct{}

func (s word) write(w io.Writer) {
	fmt.Fprintf(w, "%s", s)
}

func (s command) write(w io.Writer) {
	fmt.Fprintf(w, "%s, just %s, now!", strings.Title(s.verb), s.verb)
}

func (s greeting) write(w io.Writer) {
	fmt.Fprintf(w, "%s, %s!", strings.Title(h.verb), strings.Title(s.noun))
}

func (r runtime) write(w io.Writer) {
	fmt.Fprintf(w, "I'm running on %s with an %s CPU", runtime.GOOS, runtime.GOARCH)
}

func splitPath(path string) []string {
	c := strings.Split(path, "/")
	return c
}

func handler(w http.ResponseWriter, r *http.Request) {
	components := strings.SplitN(r.URL.Path[1:], "/", 2)
	fmt.Printf("Got request path: '%v' split into %v of len %v\n", r.URL.Path, components, len(components))
	if len(components) == 0 || len(components[0]) == 0 {
		http.NotFound(w, r)
		return
	}
	v := verb{verb: components[0]}
	if len(components) == 1 || len(components[1]) == 0 {
		v.handler(w, r)
		return
	}
	n := noun{verb: v, noun: components[1]}
	n.handler(w, r)
}

func main() {
	port := 8888
	fmt.Printf("Started server on %d\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
