package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
)

type command struct {
	verb string
}

type greeting struct {
	verb string
	noun string
}

func (c command) write(w io.Writer) {
	fmt.Fprintf(w, "%s, just %s, now!", strings.Title(c.verb), c.verb)
}

func (g greeting) write(w io.Writer) {
	fmt.Fprintf(w, "%s, %s!", strings.Title(g.verb), strings.Title(g.noun))
}

func writeRuntime(w io.Writer) {
	fmt.Fprintf(w, "I'm running on %s with an %s CPU", runtime.GOOS, runtime.GOARCH)
}

/*
	command{verb}, "/text/" -> render greeting
*/
func (c command) handlerGreeting(w http.ResponseWriter, r *http.Request) {
	g := greeting{c.verb, r.URL.Path[1 : len(r.URL.Path)-1]}
	g.write(w)
}

/*
	"/verb/" -> render verb
	"/verb/text" -> command{verb}.handlerGreeting("/text")
*/
func handlerCommand(w http.ResponseWriter, r *http.Request) {
	prefix, rest := splitPath(r.URL.Path)
	fmt.Printf("Handler command with prefix '%v' and rest '%v'\n", prefix, rest)
	c := command{prefix[1:]}
	if len(rest) == 1 {
		c.write(w)
		return
	}
	http.StripPrefix(prefix, http.HandlerFunc(c.handlerGreeting)).ServeHTTP(w, r)
}

/*
	"/" -> NotFound
	"/text" -> handlerCommand("/text")
*/
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) == 1 {
		http.NotFound(w, r)
		return
	}
	handlerCommand(w, r)
}

/*
	"text" -> redirect to "text/"
	"text/" -> handlerRoot("text/")
*/
func handlerCanonicalize(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "/") {
		r.URL.Path = r.URL.Path + "/"
		http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
	}
	handlerRoot(w, r)
}

/*
	"text" -> log, handlerCanonicalize("text")
*/
func handlerLog(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got request url '%v' with path '%v'\n", r.URL, r.URL.Path)
	handlerCanonicalize(w, r)
}

func main() {
	port := 10001
	fmt.Printf("Started server on %d\n", port)
	http.HandleFunc("/", handlerLog)
	http.Handle("/favicon.ico", http.FileServer(FS(false)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

/*
	"/" -> "", "/"
	"/text/" -> "/text", "/"
	"/prefix/text/" ->"/prefix", "/text/"
*/
func splitPath(url string) (first string, rest string) {
	i := strings.Index(url[1:], "/")
	if i < 0 {
		return "", url
	}
	return url[:i+1], url[i+1:]
}
