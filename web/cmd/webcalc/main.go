package main

import (
	_ "expvar"
	"os"

	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/alexgula/go-playground/web/path"
)

type Middleware func(http.Handler) http.Handler

/*
	"/x/y" -> render x / y
*/
func handlerDiv(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path[1:], "/")
	log.Printf("Handle division with parts %#v\n", parts)

	x, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Fprintf(w, "X is not float64: %v", err)
		return
	}

	y, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		fmt.Fprintf(w, "Y is not float64: %v", err)
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
	op, url := path.PopPrefix(r.URL.Path)
	log.Printf("Handle operation with parts %q and %q\n", op, url)

	if op == "/div" {
		http.StripPrefix(op, http.HandlerFunc(handlerDiv)).ServeHTTP(w, r)
	} else if op == "/panic" {
		http.StripPrefix(op, http.HandlerFunc(handlerPanic)).ServeHTTP(w, r)
	} else {
		fmt.Fprintf(w, "Unrecognized command: %q", op)
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
func normalizer(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/") {
			r.URL.Path = r.URL.Path + "/"
			http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
			return
		}
		h.ServeHTTP(w, r)
	}
}

/*
	"text" -> log, handlerCanonicalize("text")
*/
func logger(l *log.Logger, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Printf("Got request url '%v' with path '%v'\n", r.URL, r.URL.Path)
		h.ServeHTTP(w, r)
	}
}

/*
	"text" -> log, handlerCanonicalize("text")
*/
func recoverer(l *log.Logger, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				l.Printf("Recovered in recoverer %#v", r)
				http.Error(w, fmt.Sprintf("%#v", r), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	}
}

func main() {
	port := 4000
	l := log.New(os.Stderr, "", log.LstdFlags)
	log.Printf("Started server on %d\n", port)
	var h http.Handler
	h = http.HandlerFunc(handlerRoot)
	h = normalizer(h)
	h = logger(l, h)
	h = recoverer(l, h)
	http.Handle("/", h)
	http.Handle("/favicon.ico", http.FileServer(FS(false)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
