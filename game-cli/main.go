package main

import (
	"github.com/faiface/mainthread"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"golang.org/x/image/colornames"
)

type (
	Game struct {
		Root *Container
	}

	baseNode struct {
		id uint64
	}

	Node interface {
		ID() uint64
	}

	Container struct {
		Active map[uint64]Active
		Visual map[uint64]Visual
	}

	RenderContext struct {
		Transform pixel.Matrix
		Target    pixel.Target
	}

	Active interface {
		Node
		Update(win *pixelgl.Window) error
	}

	Visual interface {
		Node
		Render(r *RenderContext)
	}
)

var (
	nodeCounter counter
)

func newBaseNode() baseNode {
	return baseNode{id: nodeCounter.next()}
}

func (b baseNode) ID() uint64 { return b.id }

func run() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		VSync:     true,
		Title:     "Hello world",
		Bounds:    pixel.R(-400, -300, 400, 300),
		Resizable: true,
	})
	if err != nil {
		panic(err)
	}
	game := Game{
		Root: NewContainer(),
	}
	en := NewExitNode()
	game.Root.Add(en)
	game.Root.Add(NewLines())
	fb := pixelgl.NewCanvas(pixel.R(-25, -25, 25, 25))
	imd := imdraw.New(nil)
	imd.Color = colornames.Salmon
	imd.Push(pixel.V(0, 0))
	imd.Circle(5, 0)
	imd.Color = colornames.Red
	imd.Push(pixel.V(-25, 0))
	imd.Push(pixel.V(25, 0))
	imd.EndShape = imdraw.SharpEndShape
	imd.Line(1)
	imd.Color = colornames.Green
	imd.Push(pixel.V(0, -25))
	imd.Push(pixel.V(0, 25))
	imd.Line(1)
	imd.Draw(fb)
	for {
		win.Clear(colornames.Burlywood)
		for _, v := range game.Root.Active {
			v.Update(win)
		}
		ctx := RenderContext{
			Transform: pixel.IM,
			Target:    win,
		}
		fb.Draw(win, pixel.IM.Moved(win.MousePosition()))
		fb.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		for _, v := range game.Root.Visual {
			v.Render(&ctx)
		}
		win.SwapBuffers()

		if en.Quit {
			return
		}
		win.WaitInput(0)
	}
}

func main() {
	mainthread.Run(run)
}
