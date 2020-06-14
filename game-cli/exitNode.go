package main

import "github.com/faiface/pixel/pixelgl"

type (
	ExitNode struct {
		baseNode
		Quit bool
	}
)

func NewExitNode() *ExitNode {
	return &ExitNode{
		baseNode: newBaseNode(),
		Quit:     false,
	}
}

func (e *ExitNode) Update(win *pixelgl.Window) error {
	e.Quit = e.Quit ||
		win.Typed() == "q" ||
		win.JustPressed(pixelgl.KeyEscape)
	return nil
}
