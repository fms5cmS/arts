# Timing out & Moving on

原文链接：[Go Concurrency Patterns: Timing out, moving on](https://go.dev/blog/concurrency-timeouts)

[timing out 示例](./timing_out_test.go)

[moving on 示例](./moving_on_test.go)

# Pipelines & Cancellation

原文链接：[Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)

Go 的并发原语使构建流数据管道变得容易，从而有效地利用 I/O 和多个 CPU。

pipeline 是一系列通过 chan 相关联的数据处理阶段，每个阶段是一组执行相同函数的 goroutine。每个阶段中的 goroutine：

- 通过 `data <-chan` 从上游接收数据
- 对收到的数据做处理，通常还会生产出新的数据
- 通过 `chan<- data` 向下游发送数据

除了第一和最后阶段只有单个输入或输出数据的 chan 外，每个阶段都可能会有多个用于接收和发送数据的 chan。

## 简单的 pipeline

[pipeline 示例](./squaring_numbers_test.go) 

## fan out fan in

Fan out：多个函数从同一个 chan 获取数据消费。这提供了一种在一组 workers 间分配工作以并行化 CPU 使用和 I/O 的方法。

Fan in：一个函数从多个 chan 读取数据复用到单个 chan 上。该功能可以从多个输入中读取并继续执行

[Fan Out Fan in 示例](./fan_out_fan_in_test.go) 

## 提前结束！

在实际的 pipeline 中，通常并不需要接收所有的值。通常会这样设计：接收者可能只需要所有值的一个子集来进行下一步处理。更常见的是，整个 pipeline 流程因为某个阶段收到了一个 error
而提前退出。任何一种情况下，接收者都不应该等待剩余的值到达，更希望可以在早期就停止生产 pipeline 后期阶段不需要的值。

如果某个阶段在消费全部数据时失败了，则发送方不断尝试发送数据会导致 goroutine 一直阻塞，从而导致泄露。

为了在下游无法处理所有数据时不会发生 goroutine 泄露，需要在这种时候控制上游退出。

一种方式是使用 [buffered chan](./buffered_chan_bad_test.go)，这种方式比较脆弱，需要固定 buffer 的长度，一旦上下游发送或接受的数量变更，还是会遇到 goroutine 泄露；

所以，我们需要在下游提供一种方法来通知上游停止发送数据了。[cancellation_test](./cancellation_test.go) 使用一个 done 的 chan 来通知上游停止发送数据，上游则是使用 select 操作来处理。该方法存在一个问题，每个下游接收者都需要知道可能被阻塞的上游发送者的数量，并控制在早期返回时向这些发送者发出信号。而跟踪这些计数较为繁琐且易出错。

我们需要一种可以通知未知数量的 goroutine 停止向下游发送数据，在 Go 种可以通过关闭 chan 的操作来完成，因为对一个已 close chan 的接受操作总是可以立即返回，并得到一个元素类型零值。见 [close_for_cancellation](./close_for_cancellation_test.go)。

pipeline 模式通过**确保所有发送的值有足够的缓冲区**或**在接收者可能放弃通道时显式地向发送者发出信号**来解除对发送者的阻塞。

注意发送操作完成后一定要关闭 chan！

## 实际示例

MD5 是一种消息摘要算法，常用于作为文件的 checksum。`md5sum filepath` 命令可以输出文件的 md5 值。

我们的示例程序与 md5sum 类似，但将单个目录作为参数，并打印该目录下每个常规文件的摘要值，按路径名排序。

非并发模式：[md5_test](./md5_test.go) 了解基本需求。

并发模式：[md5_parallel_test](./md5_parallel_test.go)，将非并发模式中的流程拆分成了 pipeline 的两个阶段（sumFiles 遍历目录下的文件，并发计算 md5 值，并将结果发送到 chan；md5All 从 chan 接收数据并处理，如果遇到 error 就提前返回）。

上面并发模式中，每个文件会开启一个 goroutine，如果目录下的文件数量很大，分配的内存可能会超过机器的可用内存！可以通过**限制并行读取的文件数量**来限制这些分配。这样就可以拆分成 pipeline 的三个阶段（遍历目录；读取和计算文件 md5 值；收集 md5 值），见 [md5_bounded_parallel_test](./md5_bounded_parallel_test.go)。
