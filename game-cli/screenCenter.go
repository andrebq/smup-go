package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type (
	screenCenter struct {
		baseNode
		gizmo  Gizmo
		center pixel.Vec
	}
)

func NewScreenCenter(gizmo Gizmo) ActiveVisual {
	return &screenCenter{
		gizmo:    gizmo,
		baseNode: newBaseNode(),
	}
}

func (s *screenCenter) Render(rc *RenderContext) {
	s.gizmo.Draw(rc.Target, pixel.IM.Moved(s.center))
}

func (s *screenCenter) Update(win *pixelgl.Window) error {
	s.center = win.Bounds().Center()
	return nil
}
