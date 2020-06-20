package main

import (
	"sync"

	"github.com/andrebq/smup-go/game-cli/helper"
	"github.com/andrebq/smup-go/game-cli/theme"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/jakecoffman/cp"
)

type (
	player struct {
		baseNode

		runInit sync.Once
		sprite  *pixel.Sprite
		pos     pixel.Vec
		speed   pixel.Vec
		shape   pixel.Rect

		invalidBody helper.DirtyFlag
		body        *cp.Body
	}
)

// NewPlayer returns an entity which is controlled by the player
func NewPlayer(initialPos pixel.Vec) Node {
	return &player{
		baseNode:    newBaseNode(),
		speed:       pixel.V(100, 100),
		pos:         initialPos,
		shape:       pixel.R(-8, 0, 8, 32),
		invalidBody: true,
	}
}

func (p *player) Update(uc *UpdateContext) error {
	p.runInit.Do(p.init)
	if p.invalidBody.Refresh() {
		p.computeBody(uc)
	}
	p.pos = helper.ToVec(p.body.Position())
	var dir pixel.Vec
	if uc.Window.Pressed(pixelgl.KeyLeft) {
		dir.X = -1
	}
	if uc.Window.Pressed(pixelgl.KeyRight) {
		dir.X = 1
	}
	p.body.ApplyImpulseAtLocalPoint(helper.ToVector(dir.Scaled(100)), cp.Vector{})
	return nil
}

func (p *player) Render(rc *RenderContext) {
	p.sprite.Draw(rc.Target, rc.Transform.Moved(p.pos))
}

func (p *player) computeBody(uc *UpdateContext) {
	body := uc.Space.AddBody(cp.NewBody(1, cp.INFINITY))
	body.SetVelocityUpdateFunc(func(body *cp.Body, gravity cp.Vector, damping, dt float64) {
		cp.BodyUpdateVelocity(body, gravity, 0.9, dt)
	})
	body.SetPosition(helper.ToVector(p.pos))

	shape := uc.Space.AddShape(cp.NewBox(body, p.shape.W(), p.shape.H(), 0))
	shape.SetFriction(0.7)
	p.body = body
}

func (p *player) init() {
	r := p.shape
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
