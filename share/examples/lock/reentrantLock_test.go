package lock

import (
	"fmt"
	"sync"
	"testing"
)

func foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func TestReentrantLockErr(t *testing.T) {
	l := &sync.Mutex{}
	foo(l)
}
