package main

import (
	"github.com/faiface/pixel"
)

type (
	mousePosition struct {
		baseNode
		gizmo Gizmo
		pos   pixel.Vec
	}
)

func NewMousePosition(gizmo Gizmo) ActiveVisual {
	return &mousePosition{
		baseNode: newBaseNode(),
		gizmo:    gizmo,
	}
}

func (mp *mousePosition) Render(rc *RenderContext) {
	mp.gizmo.Draw(rc.Target, pixel.IM.Moved(mp.pos))
}

func (mp *mousePosition) Update(uc *UpdateContext) error {
	mp.pos = uc.Window.MousePosition()
	return nil
}
