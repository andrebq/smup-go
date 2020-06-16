package main

type (
	origin struct {
		gizmo Gizmo
		baseNode
	}
)

func NewOrigin(g Gizmo) Visual {
	return &origin{
		gizmo:    g,
		baseNode: newBaseNode(),
	}
}

func (s *origin) Render(rc *RenderContext) {
	s.gizmo.Draw(rc.Target, rc.Transform)
}
