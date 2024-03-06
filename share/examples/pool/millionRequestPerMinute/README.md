
这里的代码是简化后的，主要使用了几个重要的结构：

- `var JobQueue chan Job` 一个全局的待处理任务队列，所有生产者会将任务发送到该队列中
    - 代码中没有初始化逻辑，不过会是一个 buffered channel
- `chan chan Job` 这个类型其实就是 worker pool，该结构的初始化见 dispatcher 相关
- `Worker` 实际执行 Job 的 worker
    - 字段 `WorkerPool chan chan Job` 该 worker 所属的 worker pool
    - 字段 `JobChannel chan Job` 该 worker 所要处理的任务队列
    - 字段 `quit chan struct{}` 用来监听退出信号
    - 方法 `Start()` 起一个新的 goroutine，在死循环中进行接下来的操作：
        - 首先会将自己注册进 worker pool
        - 然后监听 JobChannel 和 quit，分别进行任务逻辑和退出操作
        - 注意：每个 worker 是一个单独的 goroutine，只会启动一次，通过死循环从 JobChannel 中消费 Job 然后机型处理
- `Dispatcher` 任务分发器
    - 初始化分发器 Dispatcher 时会根据设置的 maxWorkers 对 worker pool 进行初始化
    - `Run()`
        - 初始化 worker pool 的所有 workers
            - 然后调用每个 worker 的 `Start()` 去监听分发过来的任务和退出信号
        - 起一个新的 goroutine ，死循环来分发 JobQueue 中的任务