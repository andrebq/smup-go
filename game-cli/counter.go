package main

import "sync/atomic"

type (
	counter struct {
		val uint64
	}
)

func (c *counter) next() uint64 {
	return atomic.AddUint64(&c.val, 1)
}
