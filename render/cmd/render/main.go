package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os/exec"

	"github.com/alexgula/go-playground/render"
)

func show(filename string) {
	out, err := exec.Command("cmd", "/C", "start", filename).CombinedOutput()
	fmt.Printf("%s", out)
	assert("Could not run command: ", err)
}

func main() {
	m := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	for x := uint8(0); x < 100; x++ {
		for y := uint8(0); y < 100; y++ {
			render.Line(m, int(x), int(y), 50, 50, color.RGBA{x + y, x - y, x & y, x})
		}
	}
	err := render.Save(m, "result.png")
	assert("Could not save result image: ", err)
	show("result.png")
}

func assert(msg string, err error) {
	if err != nil {
		log.Fatal(fmt.Sprint(msg, err))
	}
}
