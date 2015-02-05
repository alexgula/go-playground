package main

import (
	"github.com/alexgula/go-playground/channel"
	"github.com/alexgula/go-playground/text/template01/model"
	"github.com/alexgula/go-playground/timeit"
	"os"
	"path/filepath"
	"text/template"
)

type NullWriter struct{}

func (this NullWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func main() {
	data := map[string]interface{}{
		"totalMessage": 1,
		"u":            model.User{"Oleksandr Gula", "alexgula@gmail.com", "My signature"},
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	tmpl := template.New("home").Funcs(template.FuncMap{
		"irange": channel.Range,
	})
	tmpl = template.Must(tmpl.ParseGlob(filepath.Join(dir, "tpl", "helper", "*.gohtml")))
	tmpl = template.Must(tmpl.ParseFiles(filepath.Join(dir, "tpl", "home.gohtml")))

	times := 10000
	f := func() {
		err = tmpl.ExecuteTemplate(NullWriter{}, "home", data)
		if err != nil {
			panic(err)
		}
	}
	timeit.RunFmt(f, times)
}
