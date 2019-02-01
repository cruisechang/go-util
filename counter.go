package util

import (
	"sync"
)

// Counter is a counting usage struct and provides functions to help couting.
type Counter struct {
	counter int
	mux     *sync.Mutex
}

// NewCounter returns a Counter instance and the input num is the init number of the counter.
func NewCounter(num int) *Counter {
	return &Counter{
		counter: num,
		mux:     &sync.Mutex{},
	}
}

// PlusPlus plus one to the counter.
func (c *Counter) PlusPlus() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.counter++
}

// SubSub sub one to the counter.
func (c *Counter) SubSub() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.counter--
}

// PlusNum plus the input number to the counter.
func (c *Counter) PlusNum(num int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.counter += num
}

// SubNum sub the input number to the counter.
func (c *Counter) SubNum(num int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.counter -= num
}

// Get returns the current count number.
func (c *Counter) Get() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.counter
}
