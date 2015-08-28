package main

import (
	_ "expvar"

	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*
	"/x/y" -> render x / y
*/
func handlerDiv(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path[1:], "/")
	log.Printf("Handle division with parts %#v\n", parts)

	x, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Fprintln(w, "X is not float64")
		return
	}

	y, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		fmt.Fprintln(w, "Y is not float64")
		return
	}

	fmt.Fprintf(w, "%v / %v = %v\n", x, y, x/y)
}

/*
	"/" -> just panic
*/
func handlerPanic(w http.ResponseWriter, r *http.Request) {
	panic([]string{"here is example of panic in HTTP handler"})
}

/*
	"/op/..." -> render operation
*/
func handlerOp(w http.ResponseWriter, r *http.Request) {
	op, url := popPrefix(r.URL.Path)
	log.Printf("Handle operation with parts %q and %q\n", op, url)

	if op == "/div" {
		http.StripPrefix(op, http.HandlerFunc(handlerDiv)).ServeHTTP(w, r)
	} else if op == "/panic" {
		http.StripPrefix(op, http.HandlerFunc(handlerPanic)).ServeHTTP(w, r)
	} else {
		fmt.Fprintln(w, "unrecognized command")
	}
}

/*
	"/" -> NotFound
	"/text" -> handlerOp("/text")
*/
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) == 1 {
		http.NotFound(w, r)
		return
	}
	handlerOp(w, r)
}

/*
	"text" -> redirect to "text/"
	"text/" -> handlerRoot("text/")
*/
func normalizer(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/") {
			r.URL.Path = r.URL.Path + "/"
			http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	}
}

/*
	"text" -> log, handlerCanonicalize("text")
*/
func logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got request url '%v' with path '%v'\n", r.URL, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

/*
	"text" -> log, handlerCanonicalize("text")
*/
func recoverer(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered in recoverer %#v", r)
				http.Error(w, fmt.Sprintf("%#v", r), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}
}

func main() {
	port := 4000
	log.Printf("Started server on %d\n", port)
	var h http.Handler
	h = http.HandlerFunc(handlerRoot)
	h = normalizer(h)
	h = logger(h)
	h = recoverer(h)
	http.Handle("/", h)
	http.Handle("/favicon.ico", http.FileServer(FS(false)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

/*
	"xxx" -> "xxx", "" (although this is impossible since all URLs start with /)
	"/" -> "/", ""
	"/text/" -> "/text", "/"
	"/prefix/text/" ->"/prefix", "/text/"
*/
func popPrefix(url string) (first string, rest string) {
	if url[0] == '/' {
		url = url[1:]
	}
	i := strings.Index(url, "/")
	if i < 0 {
		return url, ""
	}
	return url[:i+1], url[i+1:]
}
