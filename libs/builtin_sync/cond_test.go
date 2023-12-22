package builtin_sync

import (
	"golang.org/x/exp/rand"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	var ready int
	total := 10
	for i := 0; i < total; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 更改等待条件，需要加锁
			c.L.Lock()
			ready++
			c.L.Unlock()

			t.Logf("运动员#%d 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}(i)
	}

	// 观察等待条件，需要加锁
	c.L.Lock()
	for ready != total {
		c.Wait() // 调用该方法前必须加锁，且每次观察条件时都需要调用一次
		t.Log("裁判员被唤醒一次")
	}
	c.L.Unlock()

	// 所有的运动员是否就绪
	t.Log("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}
