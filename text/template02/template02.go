package main

import (
	"github.com/alexgula/go-playground/text/template02/model"
	"github.com/alexgula/go-playground/text/template02/tpl"
	"github.com/alexgula/go-playground/timeit"
)

func main() {
	user := model.User{"Oleksandr Gula", "alexgula@gmail.com", "My signature"}
	times := 10000
	f := func() {
		_ = tpl.Home(1, &user)
	}
	timeit.RunFmt(f, times)
}
