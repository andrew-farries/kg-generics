package billable

import (
	"sync"
)

type Channel[T any] struct {
	ch         chan T
	sendsMu    sync.RWMutex
	sends      int
	receivesMu sync.RWMutex
	receives   int
}

func NewChannel[T any](capacity int) Channel[T] {
	return Channel[T]{
		ch: make(chan T, capacity),
	}
}

func (c *Channel[T]) Send(x T) {
	c.ch <- x

	c.sendsMu.Lock()
	defer c.sendsMu.Unlock()
	c.sends += 1
}

func (c *Channel[T]) Receive() T {
	x := <-c.ch

	c.receivesMu.Lock()
	defer c.receivesMu.Unlock()
	c.receives += 1
	return x
}

func (c *Channel[T]) Sends() int {
	c.sendsMu.RLock()
	defer c.sendsMu.RUnlock()

	return c.sends
}

func (c *Channel[T]) Receives() int {
	c.receivesMu.RLock()
	defer c.receivesMu.RUnlock()

	return c.receives
}
