package main

import (
	"fmt"
	"image/color"
	"sync"

	"github.com/andrebq/smup-go/game-cli/helper"
	"github.com/andrebq/smup-go/game-cli/theme"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/jakecoffman/cp"
)

type (
	WorldMap struct {
		baseNode

		mapData   map[cellID]cell
		bodies    map[cellID]*cp.Body
		tileCache map[cellMaterial]*pixel.Sprite

		subtileSize int16

		currentCellSprite *pixel.Sprite
		currentCell       cellID

		rebuildPhysics helper.DirtyFlag

		initLock sync.Once
	}

	cell struct {
		material cellMaterial
		hasBody  bool
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
		baseNode:    newBaseNode(),
		mapData:     make(map[cellID]cell),
		tileCache:   make(map[cellMaterial]*pixel.Sprite),
		bodies:      make(map[cellID]*cp.Body),
		subtileSize: 16,
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

func (w *WorldMap) Update(uc *UpdateContext) error {
	w.initLock.Do(w.init)
	if uc.Window.Pressed(pixelgl.MouseButton1) {
		pos := uc.Window.MousePosition()
		cellID := w.pixToCellID(pos)
		w.mapData[cellID] = cell{material: solidMaterial}
		w.rebuildPhysics = true
	}
	w.currentCell = w.pixToCellID(uc.Window.MousePosition())
	if w.rebuildPhysics.Refresh() {
		w.refresh(uc)
	}
	return nil
}

func (w *WorldMap) refresh(uc *UpdateContext) {
	// this way of keeping the physics world is toooooooooooo lazy and reallly slow
	// but I'm just testing stuff
	//
	// a better way would be to use chunks and only recompute the chunk which was updated
	for k, v := range w.mapData {
		if !v.hasBody {
			w.bodies[k] = w.computeBodyFor(k, uc)
			v.hasBody = true
			w.mapData[k] = v
		}
	}
}

func (w *WorldMap) computeBodyFor(c cellID, uc *UpdateContext) *cp.Body {
	bounds := w.cellIDToBounds(c)
	body := uc.Space.AddBody(cp.NewBody(1, cp.INFINITY))
	body.SetPosition(helper.ToVector(bounds.Center()))
	body.SetType(cp.BODY_STATIC)

	uc.Space.AddShape(cp.NewBox(body, bounds.W(), bounds.H(), 0))

	return body
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
	w.rebuildPhysics = true
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
	im.Color = theme.Darken(c, 0.5)
	im.Push(r.Center())
	im.Circle(3, 0)
	im.Draw(cv)
	return pixel.NewSprite(cv, r)
}

func (w *WorldMap) pixToCellID(pos pixel.Vec) cellID {
	ix := int16(int16(pos.X) / w.subtileSize)
	iy := int16(int16(pos.Y) / w.subtileSize)
	return cellID{row: iy, col: ix}
}

func (w *WorldMap) cellIDToPix(cell cellID) pixel.Vec {
	sz := w.subtileSize
	halfSz := sz / 2
	return pixel.V(float64(cell.col*sz+halfSz), float64(cell.row*sz+halfSz))
}

func (w *WorldMap) cellIDToBounds(cell cellID) pixel.Rect {
	center := w.cellIDToPix(cell)
	halfSz := float64(w.subtileSize / 2)
	return pixel.R(center.X-halfSz, center.Y-halfSz, center.X+halfSz, center.Y+halfSz)
}

func (c cellID) String() string {
	return fmt.Sprintf("[%v,%v]", c.row, c.col)
}
