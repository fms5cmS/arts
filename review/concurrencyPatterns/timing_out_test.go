package concurrencyPatterns

import (
	"testing"
	"time"
)

func TestTimingOut(t *testing.T) {
	ch := make(chan int)
	select {
	case <-ch:
		// todo something
	case <-time.After(time.Second):
		// timeout, then return
	}
}
