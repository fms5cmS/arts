package concurrencyPatterns

import (
	"fmt"
	"sync"
	"testing"
)

func merge(cs ...<-chan int) <-chan int {
	// 用于确保在所有数据发送完成后再 close chan，否则向已关闭的 chan 发送数据会 panic
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(cs))
	// 合并数据
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
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

func TestFanOutFanIn(t *testing.T) {
	c1 := sq(gen(2, 4, 6, 8, 10, 12, 14, 16, 18, 20))
	c2 := sq(gen(1, 3, 5, 7, 9, 11, 13, 15, 17, 19))
	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}
