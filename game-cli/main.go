package main

import (
	"github.com/faiface/mainthread"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"golang.org/x/image/colornames"
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
	fb := pixelgl.NewCanvas(pixel.R(-25, -25, 25, 25))
	imd := imdraw.New(nil)
	imd.Color = colornames.Aquamarine
	imd.Push(pixel.V(0, 0))
	imd.Circle(25, 0)
	imd.Draw(fb)
	for {
		win.Clear(colornames.Yellowgreen)
		win.UpdateInput()
		if win.Typed() == "q" {
			return
		}
		fb.Draw(win, pixel.IM.Moved(win.MousePosition()))
		win.SwapBuffers()
	}
}

func main() {
	mainthread.Run(run)
}
