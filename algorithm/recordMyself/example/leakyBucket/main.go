package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)

func main() {
	// 设置 RPS(Requests Per Second)，即桶容量
	rl := ratelimit.New(100)
	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
