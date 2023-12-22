package lock

import (
	"sync"
	"testing"
	"time"
)

type CounterRW struct {
	mu    sync.RWMutex
	count uint64
}

func (c *CounterRW) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *CounterRW) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func TestRWMutex_InStruct(t *testing.T) {
	var counter CounterRW
	for i := 0; i < 10; i++ { // 10 个 reader
		go func() {
			for {
				t.Log(counter.Count())
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for { // 一个 writer
		counter.Incr()
		time.Sleep(time.Second)
	}
}
