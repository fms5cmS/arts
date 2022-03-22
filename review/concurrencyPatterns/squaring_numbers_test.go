package concurrencyPatterns

import (
	"fmt"
	"testing"
)

// 使用 unbuffered chan，如果发送多个数据的话，必须单独开 goroutine ！

// 用于生产一系列整数并向 chan 发送
func gen(nums ...int) <-chan int {
	out := make(chan int)
	// 单独开启 goroutine 向下游发送数据
	// 这样就可以：这里向 out 发送一个数据，下游从 out 接收一个数据来处理，上游和下游一起运行
	go func() {
		// 注意：这里使用的是 unbuffered chan，下游如果处理数据失败导致报错退出
		// chan 的发送方会因为接收方未准备好而一直阻塞，导致 goroutine 泄露
		for _, n := range nums {
			out <- n
		}
		close(out) // 数据发送完毕后，关闭 chan
	}()
	return out
}

// 处理从上游接收到的数据（求平方）然后向下游发送数据
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() { // 单独开启 goroutine 向下游发送数据
		for n := range in {
			out <- n * n
		}
		close(out) // 数据发送完毕后，关闭 chan
	}()
	return out
}

func TestPipeline(t *testing.T) {
	for n := range sq(sq(gen(2, 3, 4, 5, 6))) {
		fmt.Println(n) // 16 then 81
	}
}
