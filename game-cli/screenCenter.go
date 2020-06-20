package main

import (
	"github.com/faiface/pixel"
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

func (s *screenCenter) Update(uc *UpdateContext) error {
	s.center = uc.Window.Bounds().Center()
	return nil
}
