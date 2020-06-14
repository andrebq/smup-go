package main

func NewContainer() *Container {
	return &Container{
		Active: make(map[uint64]Active),
		Visual: make(map[uint64]Visual),
	}
}

func (c *Container) Add(a Node) *Container {
	if a.ID() == 0 {
		if a, ok := a.(interface{ SetID(uint64) }); ok {
			a.SetID(nodeCounter.next())
		}
	}
	c.addActive(a)
	c.addVisual(a)
	return c
}

func (c *Container) addActive(a Node) {
	if a, ok := a.(Active); ok {
		c.Active[a.ID()] = a
	}
}

func (c *Container) addVisual(a Node) {
	if a, ok := a.(Visual); ok {
		c.Visual[a.ID()] = a
	}
}
