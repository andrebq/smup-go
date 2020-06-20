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

func (e *ExitNode) Update(uc *UpdateContext) error {
	win := uc.Window
	e.Quit = e.Quit ||
		win.Typed() == "q" ||
		win.JustPressed(pixelgl.KeyEscape)
	return nil
}
