package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"os/exec"
)

func dot() {
	im := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	im.Set(50, 50, color.RGBA{255, 0, 0, 255})
	f, err := os.Create("result.png")
	defer f.Close()
	assert(err, "Could not create file: ")
	err = png.Encode(f, im)
	assert(err, "Could not write image: ")
}

func main() {
	dot()
	out, err := exec.Command("cmd", "/C", "start", "result.png").CombinedOutput()
	fmt.Printf("%s", out)
	assert(err, "Could not run command: ")
}

func assert(err error, msg string) {
	if err != nil {
		log.Fatal(fmt.Sprint(msg, err))
	}
}
