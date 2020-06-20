package helper

import (
	"github.com/faiface/pixel"
	"github.com/jakecoffman/cp"
)

// V is a short cut to construct a Chipmunk Physics Vector
func V(x, y float64) cp.Vector {
	return cp.Vector{X: x, Y: y}
}

// ToVector takes a pixel.Vec and returns the cp.Vector most likely the compiler
// will inline this call
func ToVector(pv pixel.Vec) cp.Vector {
	return cp.Vector{
		X: pv.X,
		Y: pv.Y,
	}
}

// ToVec is the opposite of ToVector
func ToVec(v cp.Vector) pixel.Vec {
	return pixel.V(v.X, v.Y)
}
