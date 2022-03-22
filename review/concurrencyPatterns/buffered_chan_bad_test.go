package concurrencyPatterns

import (
	"fmt"
	"sync"
	"testing"
)

/** pipeline 的下游某个阶段处理数据失败时，上游的发送操作会一直阻塞导致 goroutine 泄露，可以使用 buffered chan，修改如下：**/

func gen2(nums ...int) <-chan int {
	// 这里创建 chan 时就可以知道要发送的数量
	out := make(chan int, len(nums))
	// 将 nums 的值复制到 chan 的缓冲区，还可以避免创建新的 goroutine
	for _, n := range nums {
		out <- n
	}
	close(out) // 数据发送完毕后，关闭 chan
	return out
}

func merge2(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	// ! merge 函数每次接收以及下游每次处理的数量为 1，所以这里设置 buffer 长度为 1
	// if we pass an additional value to gen, or if the downstream stage reads any fewer values, we will again have blocked goroutines.
	// 这个操作是脆弱的，一旦 merge 接收和下游处理数量发生变化，又会再次出现 goroutine 阻塞泄露
	out := make(chan int, 1)
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
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func TestBuffered(t *testing.T) {
	c1 := sq(gen2(2, 4, 6, 8, 10, 12, 14, 16, 18, 20))
	c2 := sq(gen2(1, 3, 5, 7, 9, 11, 13, 15, 17, 19))
	for n := range merge2(c1, c2) {
		fmt.Println(n)
	}
}
