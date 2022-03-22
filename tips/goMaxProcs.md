在 fasthttp 和 ants 库的代码中，针对 GOMAXPROCS 会有一个特殊的性能优化：

```go
var workerChanCap = func() int {
	// Use blocking workerChan if GOMAXPROCS=1.
	// This immediately switches Serve to WorkerFunc, which results
	// in higher performance (under go1.5 at least).
	if runtime.GOMAXPROCS(0) == 1 {
		return 0
	}

	// Use non-blocking workerChan if GOMAXPROCS>1,
	// since otherwise the Serve caller (Acceptor) may lag accepting
	// new connections if WorkerFunc is CPU-bound.
	return 1
}()
```

当 GOMAXPROCS == 1 时，channel 的缓冲区长度会设置为 0！这样的话，当发送方向 channel 发送时会直接阻塞住，然后执行流程立刻切换到接收方 goroutine。

当 GOMAXPROCS > 1 时，如果接收方的操作是 CPU 密集型的，使用带缓冲区的 channel 可以防止发送方阻塞。

注意：上面的 workerChanCap 是一个 int 类型的值！因为函数已经执行了。
