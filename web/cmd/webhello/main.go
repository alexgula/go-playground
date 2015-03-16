package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

type verb struct {
	verb string
}

type noun struct {
	verb
	noun string
}

func (h verb) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>%s, just %s!</p>", strings.Title(h.verb), h.verb)
	fmt.Printf("%s, just %s!\n", strings.Title(h.verb), h.verb)
	fmt.Fprintf(w, "<p>I'm running on %s with an %s CPU</p>", runtime.GOOS, runtime.GOARCH)
}

func (h noun) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>%s, %s!</p>", strings.Title(h.verb.verb), h.noun)
	fmt.Printf("%s, %s!\n", strings.Title(h.verb.verb), h.noun)
	fmt.Fprintf(w, "<p>I'm running on %s with an %s CPU</p>", runtime.GOOS, runtime.GOARCH)
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
