package builtin_sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	var pool any
	var once sync.Once
	var initFn = func() {
		fmt.Println("run...")
		pool = 1
	}
	for i := 0; i < 10; i++ {
		once.Do(initFn)
		t.Log(pool)
	}
}

func TestOnceValue(t *testing.T) {
	var initFn = func() any {
		fmt.Println("run...")
		return 1
	}
	var poolGenerator = sync.OnceValue(initFn)
	for i := 0; i < 10; i++ {
		t.Log(poolGenerator())
	}
}
