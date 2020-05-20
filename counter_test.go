package goutil

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var wg sync.WaitGroup

	counter := NewCounter(0)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Counter) {
			defer wg.Done()
			c.PlusPlus()
		}(&wg, counter)
	}
	wg.Wait()
	if counter.Get() != 100 {
		t.Errorf("TestCounter: expect->100, actual->%d", counter.Get())
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Counter) {
			defer wg.Done()
			c.SubSub()
		}(&wg, counter)
	}
	wg.Wait()
	if counter.Get() != 0 {
		t.Errorf("TestCounter: expect->0, actual->%d", counter.Get())
	}
}
