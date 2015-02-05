package main

import (
	"github.com/alexgula/go-playground/text/template03/model"
	"github.com/alexgula/go-playground/timeit"
	"github.com/eknkc/amber"
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
	templates, err := amber.CompileDir("templates/", amber.DefaultDirOptions, amber.DefaultOptions)
	tmpl := templates["master"]

	times := 10000
	f := func() {
		err = tmpl.ExecuteTemplate(NullWriter{}, "master", data)
		if err != nil {
			panic(err)
		}
	}
	timeit.RunFmt(f, times)
}
