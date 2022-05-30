package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//limiter := NewFixedWindow(100, time.Second)

	limiter := NewSlidingWindow(time.Second, time.Second, 100)

	go func() {
		for {
			if runtime.NumGoroutine() > 110 {
				fmt.Printf("the number of goroutines is %d\n", runtime.NumGoroutine())
				panic("+++++++++++ failed ++++++++++")
			}
		}
	}()
	time.Sleep(999 * time.Millisecond)
	wg := sync.WaitGroup{}
	for i := 0; i < 200 && limiter.Allow(); i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			time.Sleep(5 * time.Millisecond)
			fmt.Printf("--> worker %d finished\n", num)
		}(i)
	}
	wg.Wait()
	fmt.Println("succeed")
}
