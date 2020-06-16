package main

import (
	"time"

	"github.com/faiface/mainthread"
	"github.com/faiface/pixel"
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

	ActiveVisual interface {
		Active
		Visual
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
	axisGizmo := NewAxisGizmo(50)
	game.Root.Add(en)
	game.Root.Add(NewOrigin(axisGizmo))
	game.Root.Add(NewMousePosition(axisGizmo))
	game.Root.Add(NewScreenCenter(axisGizmo))
	game.Root.Add(NewWorldMap())
	for {
		win.Clear(colornames.Burlywood)
		for _, v := range game.Root.Active {
			v.Update(win)
		}
		ctx := RenderContext{
			Transform: pixel.IM,
			Target:    win,
		}
		for _, v := range game.Root.Visual {
			v.Render(&ctx)
		}
		win.SwapBuffers()

		if en.Quit {
			return
		}
		win.WaitInput(time.Second / 24)
	}
}

func main() {
	mainthread.Run(run)
}
