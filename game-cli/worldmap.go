package main

import (
	"fmt"
	"image/color"
	"sync"

	"github.com/andrebq/smup-go/game-cli/theme"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type (
	WorldMap struct {
		baseNode

		mapData   map[cellID]cell
		tileCache map[cellMaterial]*pixel.Sprite

		currentCellSprite *pixel.Sprite
		currentCell       cellID

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
		tileCache: make(map[cellMaterial]*pixel.Sprite),
	}
}

func (w *WorldMap) Render(rc *RenderContext) {
	for k, v := range w.mapData {
		tile := w.tileCache[v.material]
		pixPos := w.cellIDToPix(k)
		tile.Draw(rc.Target, rc.Transform.Moved(pixPos))
	}

	w.currentCellSprite.Draw(rc.Target, rc.Transform.Moved(w.cellIDToPix(w.currentCell)))
}

func (w *WorldMap) Update(win *pixelgl.Window) error {
	w.initLock.Do(w.init)
	if win.Pressed(pixelgl.MouseButton1) {
		pos := win.MousePosition()
		cellID := w.pixToCellID(pos)
		w.mapData[cellID] = cell{material: solidMaterial}
	}
	w.currentCell = w.pixToCellID(win.MousePosition())
	return nil
}

func (w *WorldMap) init() {
	for i := 0; i < 100; i++ {
		w.mapData[cellID{
			row: 0,
			col: int16(i - 50),
		}] = cell{material: solidMaterial}
	}
	w.currentCellSprite = wireframeTile(pixel.R(0, 0, 16, 16), theme.Contour)
	w.tileCache[solidMaterial] = solidColorTile(pixel.R(0, 0, 16, 16), theme.Sand)
}

func wireframeTile(r pixel.Rect, c color.Color) *pixel.Sprite {
	cv := pixelgl.NewCanvas(r)
	im := imdraw.New(nil)
	im.Color = c
	im.Push(r.Min)
	im.Push(r.Max)
	im.Rectangle(2)
	im.Draw(cv)

	return pixel.NewSprite(cv, r)
}

func solidColorTile(r pixel.Rect, c color.Color) *pixel.Sprite {
	cv := pixelgl.NewCanvas(r)
	im := imdraw.New(nil)
	im.Color = c
	im.Push(r.Min)
	im.Push(r.Max)
	im.Rectangle(0)
	im.Draw(cv)

	return pixel.NewSprite(cv, r)
}

func (w *WorldMap) pixToCellID(pos pixel.Vec) cellID {
	ix := int16(int(pos.X) / 16)
	iy := int16(int(pos.Y) / 16)
	return cellID{row: iy, col: ix}
}

func (w *WorldMap) cellIDToPix(cell cellID) pixel.Vec {
	return pixel.V(float64(cell.col*16+8), float64(cell.row*16+8))
}

func (c cellID) String() string {
	return fmt.Sprintf("[%v,%v]", c.row, c.col)
}
