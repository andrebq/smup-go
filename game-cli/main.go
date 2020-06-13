package main

import (
	"github.com/faiface/mainthread"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func run() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		VSync:  true,
		Title:  "Hello world",
		Bounds: pixel.R(0, 0, 800, 600),
	})
	if err != nil {
		panic(err)
	}
	for {
		win.Update()
		if win.Typed() == "q" {
			return
		}
	}
}

func main() {
	mainthread.Run(run)
}
