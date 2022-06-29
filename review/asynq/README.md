[项目地址](https://github.com/hibiken/asynq)

一个可以用于对任务排队，然后由 workers 异步处理的第三方库。任务队列可以被用作一种在多个机器之间分发任务的机制（mechanism）。

特性：
- 保证每个任务至少会被执行一次！
- 失败任务重试
- worker 崩溃时可以自动恢复任务
- 优先级队列
- 因为 Redis 的写入很快，所以添加任务的延迟很低
- 可以使用 Unique Option 来避免任务的重复处理
- 能够暂停队列以停止处理队列中的任务
- 周期性任务（Periodic tasks）
- 整合了 Prometheus 来采集以及可视化队列指标
- 可以使用 Web UI 或命令行的方式来检查和远程控制队列和任务
- 可以利用 Redis 集群来做自动 sharding 及高可用，也可以利用 Redis 哨兵机制来实现高可用

