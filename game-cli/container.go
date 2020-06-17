package main

import "github.com/faiface/pixel/pixelgl"

type (
	Container struct {
		nodes map[uint64]Node

		active []uint64
		visual []uint64
	}
)

func NewContainer() *Container {
	return &Container{
		nodes: make(map[uint64]Node),
	}
}

func (c *Container) Update(win *pixelgl.Window) error {
	var firstErr error
	for _, v := range c.active {
		err := c.nodes[v].(Active).Update(win)
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (c *Container) Render(rc *RenderContext) {
	for _, v := range c.visual {
		c.nodes[v].(Visual).Render(rc)
	}
}

func (c *Container) Add(a Node) *Container {
	if a.ID() == 0 {
		if a, ok := a.(interface{ SetID(uint64) }); ok {
			a.SetID(nodeCounter.next())
		} else {
			// if we are here, there is a programmer error and it doesn't make sense to proceed
			panic("cannot add a node with ID() == 0 and without SetID(uint64) method")
		}
	}
	c.nodes[a.ID()] = a
	c.addActive(a)
	c.addVisual(a)
	return c
}

func (c *Container) addActive(a Node) {
	if a, ok := a.(Active); ok {
		c.active = append(c.active, a.ID())
	}
}

func (c *Container) addVisual(a Node) {
	if a, ok := a.(Visual); ok {
		c.visual = append(c.visual, a.ID())
	}
}
