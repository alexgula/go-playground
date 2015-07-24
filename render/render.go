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

// Canvas is image that can be written to
type Canvas interface {
	image.Image
	Set(x, y int, c color.Color)
}

func absd(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func signd(x, y int) int {
	if x > y {
		return 1
	}
	return -1
}

func line(m Canvas, x0, y0, x1, y1 int, c color.Color) Canvas {
	var swap bool
	if absd(x0, x1) < absd(y0, y1) {
		x0, y0, x1, y1 = y0, x0, y1, x1
		swap = true
	}

	if x0 > x1 {
		x0, y0, x1, y1 = x1, y1, x0, y0
	}

	dx := absd(x0, x1)
	dx2 := dx * 2
	dy := signd(y1, y0)

	de := absd(y0, y1) * 2
	e := 0

	for x, y := x0, y0; x <= x1; x++ {
		if swap {
			m.Set(y, x, c)
		} else {
			m.Set(x, y, c)
		}
		e += de
		if e > dx {
			y += dy
			e -= dx2
		}
	}
	return m
}

func lineopt(m Canvas, x0, y0, x1, y1 int, c color.Color) Canvas {
	if absd(x0, x1) > absd(y0, y1) {
		if x0 > x1 {
			x0, y0, x1, y1 = x1, y1, x0, y0
		}

		dx := absd(x0, x1)
		dx2 := dx * 2
		dy := signd(y1, y0)

		de := absd(y0, y1) * 2
		e := 0

		for x, y := x0, y0; x <= x1; x++ {
			m.Set(x, y, c)
			e += de
			if e > dx {
				y += dy
				e -= dx2
			}
		}
	} else {
		if y0 > y1 {
			x0, y0, x1, y1 = x1, y1, x0, y0
		}

		dy := absd(y0, y1)
		dy2 := dy * 2
		dx := signd(x1, x0)

		de := absd(x0, x1) * 2
		e := 0

		for x, y := x0, y0; y <= y1; y++ {
			m.Set(x, y, c)
			e += de
			if e > dy {
				x += dx
				e -= dy2
			}
		}
	}
	return m
}

func save(m image.Image, fn string) {
	f, err := os.Create(fn)
	defer f.Close()
	assert("Could not create file: ", err)
	err = png.Encode(f, m)
	assert("Could not write image: ", err)
}

func show(filename string) {
	out, err := exec.Command("cmd", "/C", "start", filename).CombinedOutput()
	fmt.Printf("%s", out)
	assert("Could not run command: ", err)
}

func main() {
	m := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	for x := uint8(0); x < 100; x++ {
		for y := uint8(0); y < 100; y++ {
			line(m, int(x), int(y), 50, 50, color.RGBA{x + y, x - y, x & y, 255})
		}
	}
	save(m, "result.png")
	show("result.png")
}

func assert(msg string, err error) {
	if err != nil {
		log.Fatal(fmt.Sprint(msg, err))
	}
}
