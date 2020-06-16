package main

import (
	"sync"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type (
	Gizmo interface {
		Draw(pixel.Target, pixel.Matrix)
	}

	axisGizmo struct {
		initLock sync.Once
		pic      *pixelgl.Canvas
		size     int
	}
)

// NewAxisGizmo returns a reusable Axis Gizmo
func NewAxisGizmo(size int) Gizmo {
	if size <= 0 {
		size = 50
	}
	return &axisGizmo{
		size: size,
	}
}

// Should be called only from a mainthread.Call
func (a *axisGizmo) Draw(target pixel.Target, im pixel.Matrix) {
	a.initLock.Do(a.init)
	a.pic.Draw(target, im)
}

func (a *axisGizmo) init() {
	halfSize := float64(a.size) / 2
	fb := pixelgl.NewCanvas(pixel.R(-halfSize, -halfSize, halfSize, halfSize))
	imd := imdraw.New(nil)
	imd.Color = colornames.Salmon
	imd.Push(pixel.V(0, 0))
	imd.Circle(5, 0)
	imd.Color = colornames.Red
	imd.Push(pixel.V(-25, 0))
	imd.Push(pixel.V(25, 0))
	imd.EndShape = imdraw.SharpEndShape
	imd.Line(1)
	imd.Color = colornames.Green
	imd.Push(pixel.V(0, -25))
	imd.Push(pixel.V(0, 25))
	imd.Line(1)
	imd.Draw(fb)
	a.pic = fb
}
