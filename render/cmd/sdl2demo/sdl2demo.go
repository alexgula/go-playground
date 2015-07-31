package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	assert("init", err)
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	assert("create window", err)
	defer window.Destroy()

	surface, err := window.GetSurface()
	assert("get surface", err)

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	err = surface.FillRect(&rect, 0xffff0000)
	assert("fill rect", err)
	err = window.UpdateSurface()
	assert("update surface", err)

	sdl.Delay(1000)
	sdl.Quit()
}

func assert(msg string, err error) {
	if err != nil {
		log.Fatal(fmt.Sprint("Error during ", msg, ": ", err))
	}
}
