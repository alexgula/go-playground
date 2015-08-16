package render

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
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

func Line(m Canvas, x0, y0, x1, y1 int, c color.Color) Canvas {
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

func Save(m image.Image, fn string) (err error) {
	f, err := os.Create(fn)
	if err != nil {
		return errors.New("Could not create a file: " + err.Error())
	}
	defer func() {
		if cerr := f.Close(); cerr != err {
			err = combineError(
				err,
				errors.New("Could not close file: "+cerr.Error()))
		}
	}()
	err = png.Encode(f, m)
	if err != nil {
		return errors.New("Could not encode PNG: " + err.Error())
	}
	return nil
}

func combineError(outer, inner error) error {
	if outer != nil {
		return errors.New(outer.Error() + " and " + inner.Error())
	}
	return inner
}
