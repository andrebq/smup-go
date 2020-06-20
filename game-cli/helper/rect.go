package helper

import (
	"github.com/faiface/pixel"
	"github.com/jakecoffman/cp"
)

// ToBB returns a bounding-box from a pixel rect
func ToBB(r pixel.Rect) cp.BB {
	return cp.BB{
		L: r.Min.X,
		B: r.Min.Y,
		R: r.Max.X,
		T: r.Max.Y,
	}
}
