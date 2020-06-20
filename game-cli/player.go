package main

import (
	"sync"

	"github.com/andrebq/smup-go/game-cli/theme"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type (
	player struct {
		baseNode

		runInit sync.Once
		sprite  *pixel.Sprite
		pos     pixel.Vec
		speed   pixel.Vec
	}
)

// NewPlayer returns an entity which is controlled by the player
func NewPlayer() Node {
	return &player{
		baseNode: newBaseNode(),
		speed:    pixel.V(1000, 1000),
	}
}

func (p *player) Update(uc *UpdateContext) error {
	p.runInit.Do(p.init)
	var dir pixel.Vec
	if uc.Window.Pressed(pixelgl.KeyLeft) {
		dir.X = -1
	}
	if uc.Window.Pressed(pixelgl.KeyRight) {
		dir.X = 1
	}
	p.pos.X += p.speed.X * dir.X * uc.Delta
	return nil
}

func (p *player) Render(rc *RenderContext) {
	p.sprite.Draw(rc.Target, rc.Transform.Moved(p.pos))
}

func (p *player) init() {
	r := pixel.R(-8, 0, 8, 32)
	pW := r.W() * 0.1
	pH := r.H() * 0.8
	cv := pixelgl.NewCanvas(r)
	im := imdraw.New(nil)
	im.Color = theme.PlayerBody
	im.Push(r.Min)
	im.Push(r.Max)
	im.Rectangle(0)
	im.Color = theme.PlayerHighlight
	im.Push(pixel.V(r.Min.X, r.Min.Y))
	im.Push(pixel.V(r.Min.X+pW, r.Min.Y+pH))
	im.Rectangle(0)
	im.Push(pixel.V(r.Min.X, r.Min.Y))
	im.Push(pixel.V(r.Min.X+pH, r.Min.Y+pW))
	im.Rectangle(0)
	im.Draw(cv)
	p.sprite = pixel.NewSprite(cv, r)
}
