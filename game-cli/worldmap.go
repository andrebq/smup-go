package main

import (
	"fmt"
	"image/color"
	"sync"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type (
	WorldMap struct {
		baseNode

		mapData   map[cellID]cell
		tileCache map[cellMaterial]*pixelgl.Canvas

		initLock sync.Once
	}

	cell struct {
		material cellMaterial
	}

	cellID struct {
		row int16
		col int16
	}

	cellMaterial byte
)

const (
	noMaterial = cellMaterial(iota)
	solidMaterial
)

func NewWorldMap() *WorldMap {
	return &WorldMap{
		baseNode:  newBaseNode(),
		mapData:   make(map[cellID]cell),
		tileCache: make(map[cellMaterial]*pixelgl.Canvas),
	}
}

func (w *WorldMap) Render(rc *RenderContext) {
}

func (w *WorldMap) Update(win *pixelgl.Window) error {
	w.initLock.Do(w.init)
	if win.JustPressed(pixelgl.MouseButton1) {
		pos := win.MousePosition()
		cellID := w.pixToCellID(pos)
		w.mapData[cellID] = cell{material: solidMaterial}
	}
	return nil
}

func (w *WorldMap) init() {
	cv := pixelgl.NewCanvas(pixel.R(0, 0, 16, 16))
	im := imdraw.New(nil)
	im.Push(pixel.V(0, 0))
	im.Push(pixel.V(16, 16))
	im.Color = color.RGBA{
		R: 238,
		G: 154,
		B: 31,
		A: 255,
	}
	im.Rectangle(0)
	im.Draw(cv)
	w.tileCache[solidMaterial] = cv
}

func (w *WorldMap) pixToCellID(pos pixel.Vec) cellID {
	ix := int16(int(pos.X) / 16)
	iy := int16(int(pos.Y) / 16)
	return cellID{row: iy, col: ix}
}

func (c cellID) String() string {
	return fmt.Sprintf("[%v,%v]", c.row, c.col)
}
