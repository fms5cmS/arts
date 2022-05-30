package main

import (
	"fmt"
	"sync"
	"time"
)

type FixedWindow struct {
	rate  int           //计数周期内最多允许的请求数
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	count int           //计数周期内累计收到的请求数
	lock  sync.Mutex
}

func (l *FixedWindow) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.count < l.rate {
		l.count++
		return true
	}
	now := time.Now()
	if now.Sub(l.begin) >= l.cycle {
		l.Reset(now)
		return true
	}
	return false
}

// Reset 重置起始时间和计数器
func (l *FixedWindow) Reset(t time.Time) {
	l.begin = t
	l.count = 0
	fmt.Println("Reset counter")
}

func NewFixedWindow(rate int, cycle time.Duration) *FixedWindow {
	return &FixedWindow{
		rate:  rate,
		begin: time.Now(),
		cycle: cycle,
		lock:  sync.Mutex{},
	}
}
