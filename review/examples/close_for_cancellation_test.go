package examples

import (
	"fmt"
	"sync"
	"testing"
)

func TestCloseForCancellation(t *testing.T) {
	// 通知上游停止发送数据
	done := make(chan struct{})
	// 关闭 done 来通知
	defer close(done)

	in := gen(2, 3)
	c1 := sq4(done, in)
	c2 := sq4(done, in)
	out := merge4(done, c1, c2)
	fmt.Println(<-out)
}

func merge4(done chan struct{}, cs ...<-chan int) <-chan int {
	// 用于确保在所有数据发送完成后再 close chan，否则向已关闭的 chan 发送数据会 panic
	var wg sync.WaitGroup
	out := make(chan int)
	for _, c := range cs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				select {
				case out <- n:
				case <-done: // 收到信号后 return，不再向下游发送数据
					return
				}
			}
		}(c)
	}
	// 会先将 out 返回，下游可以实时处理 out 中的数据，这里开 goroutine 来关闭 out，可以边 merge 输入数据，同时下游边处理数据
	// 而不必非等到 merge 完成后再处理数据
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func sq4(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}
