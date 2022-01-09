package examples

import (
	"fmt"
	"sync"
	"testing"
)

func TestCancellation(t *testing.T) {
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)
	// 用于告诉上游剩下的数据不用再发送了
	// 因为有两个发送方，所以这里 buffer 长度为 2
	done := make(chan struct{}, 2)
	out := merge3(done, c1, c2)
	fmt.Println(<-out) // 4 or 9
	done <- struct{}{}
	done <- struct{}{}
}

func merge3(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1)
	for _, c := range cs {
		wg.Add(1)
		go func(c <-chan int) {
			for n := range c {
				select {
				case out <- n:
				case <-done: // 一旦收到结束的信号就不再向下游发送数据了
				}
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
