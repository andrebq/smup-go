package main

import (
	"github.com/andrebq/smup-go/game-cli/metrics"
	"github.com/andrebq/smup-go/game-cli/theme"
	"github.com/faiface/mainthread"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-gl/glfw/v3.3/glfw"
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

	RenderContext struct {
		Transform pixel.Matrix
		Target    pixel.Target
	}

	UpdateContext struct {
		Delta  float64
		Window *pixelgl.Window
	}

	Active interface {
		Node
		Update(*UpdateContext) error
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
	delta := metrics.GetDelta()
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Undecorated: true,
		Title:       "Hello world",
		Bounds:      pixel.R(-400, -300, 400, 300),
		Resizable:   true,
	})
	win.SetCursorVisible(false)
	if err != nil {
		panic(err)
	}
	game := Game{
		Root: NewContainer(),
	}
	en := NewExitNode()
	// axisGizmo := NewAxisGizmo(50)
	game.Root.Add(en)
	game.Root.Add(NewWorldMap())
	game.Root.Add(NewPlayer())
	for {
		win.UpdateInput()
		win.Clear(theme.Base)
		uc := UpdateContext{
			Window: win,
			Delta:  delta.Tick(),
		}
		game.Root.Update(&uc)
		ctx := RenderContext{
			Transform: pixel.IM,
			Target:    win,
		}
		game.Root.Render(&ctx)
		win.SwapBuffers()
		if en.Quit || win.Closed() {
			return
		}
	}
}

func main() {
	go metrics.RunMetricsHTTP()
	mainthread.Run(run)
}
