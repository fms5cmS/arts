[项目地址](https://github.com/hibiken/asynq)
[文档](https://github.com/hibiken/asynq/wiki)

一个可以用于对任务排队，然后由 workers 异步处理的第三方库。任务队列可以被用作一种在多个机器之间分发任务的机制（mechanism）。

特性：
- 保证每个任务至少会被执行一次！
- 失败任务重试
- worker 崩溃时可以自动恢复任务
- 优先级队列，值越大优先级越高
- 因为 Redis 的写入很快，所以添加任务的延迟很低
- 可以使用 Unique Option 来避免任务的重复处理
- 能够暂停队列以停止处理队列中的任务
- 周期性任务（Periodic tasks）
- 整合了 Prometheus 来采集以及可视化队列指标
- 可以使用 [Web UI](https://github.com/hibiken/asynqmon#readme) 或命令行的方式来检查和远程控制队列和任务
- 可以利用 Redis 集群来做自动 sharding 及高可用，也可以利用 Redis 哨兵机制来实现高可用
  - `RedisClusterClientOpt`、`RedisFailoverClientOpt`
- graceful shutdown
  - 建议先发送 SIGTSTP 信号以停止处理新的任务，等所有处理中的任务结束后发送 SIGTERM 信号来结束程序
  - 如果直接发送 TERM 信号，程序会等待 8s 以便完成正在处理中的任务，8s 后未完成的任务会被重置为 pending 状态，以便程序重启后再进行处理

```shell
kill -TSTP <pid> # stop processing new tasks
kill -TERM <pid> # shutdown the server
```

---

版本：v0.23.0

**注：Redis 命令相关中，所有传入的动态值均以驼峰格式来表示！**

# Task

- `asynq.Task` 
  - typename，任务对应的任务类型
  - payload，任务执行所需数据
  - opts，任务的可选配置，用于指定任务处理的行为，如重试次数等，详见 Option
  - w，任务执行结果的输出流

- `asynq.NewTask(typename string, payload []byte, opts ...Option)` 创建任务

任务的生命周期：

- `client.Enqueue(task1, asynq.ProcessIn(24*time.Hour))`
  - 在执行时间（24h）到来之前一直处于 scheduled 状态
  - 到了 24h 后，会转变为 pending 状态，然后转为 active 状态
  - 如果执行成功，task 的数据会从 Redis 中删除
  - 如果执行失败，会转为 retry 状态等待之后重试
    - 重试间隔到达后，会再次转为 pending 状态，然后转为 active 状态
    - 重复这个步骤直到 task 执行成功，或达到重试上限
    - 如果达到重试上限依然执行失败，则转为 archived 状态
- `client.Enqueue(task2)`
  - 与 task1 唯一不同的是，task2 没有 scheduled 状态，会直接转为 pending 状态
- `client.Enqueue(task3, asynq.Retention(2*time.Hour))`
  - 与 task2 不同的是，当任务被成功处理后，任务仍会在队列中保留 2h（处于 completed 状态）

# Option

### 任务唯一性

任务的唯一性有两种方式可以实现：

方案一：由 asynq 创建一个唯一锁

- `uniqueOption` 通过 `asynq.Unique()` 来生成，可以保证在 Redis 中排队的任务只会有一个副本，当想要重复的 tasks 时，可以确保不会创建冗余的 work。

client 入队 task 时会先为 task 获取一个锁，获取成功，则 task 入队成功；如果 锁已经被另一个 task 持有，client 会返回 error。

为了避免锁被 task 一直持有，所以唯一锁还需要设置 ttl，ttl 过期、又或者在 ttl 时间段内持有锁的 task 成功完成，锁会被释放。

这种方式只是尽最大可能保证唯一，所以，如果在 task 被执行之前锁已经过期了，就有可能入队重复的 task。

task 的唯一性基于 typename、payload、以及队列名称

方案二：自己生成一个唯一的 task id

- `taskIDOption` 通过 `asynq.TaskID()` 来生成、，可以保证任何时候都只会有一个给定的 task id，如果入队一个相同 task id 的其他任务，会返回 error

```go
// First task should be ok
_, err := client.Enqueue(task, asynq.TaskID("mytaskid"))

// Second task will fail, err is ErrTaskIDConflict (assuming that the first task didn't get processed yet)
_, err = client.Enqueue(task, asynq.TaskID("mytaskid"))
```

## 指定任务执行时间

`processAtOption` 指定任务执行的时间点；
`processInOption` 指定任务在多长时间以后执行。

二者会被统一处理到 processAt 属性中 


# Client

- `asynq.NewClient(r RedisConnOpt)` 根据 redis 配置创建一个 client
  - client 中会包含一个用于访问 Redis 的 Redis client，以及一个返回当前时间的时钟 CLock
  - client 可以用于查询和修改（mutate）任务队列


## Enqueue Task

- `func (c *Client) Enqueue(task *Task, opts ...Option) (*TaskInfo, error)` 入队任务
  - 注意：`NewTask()` 配置的 options 会被 `Enqueue()` 的相同类型的配置覆盖！
  - 注意：默认情况下最大重试次数为 25，最大超时时间为 3min
  - 如果没有配置 `ProcessAt` 和 `ProcessIn` 任一 option，则任务会被立即执行

```go
func (c *Client) Enqueue(task *Task, opts ...Option) (*TaskInfo, error) {
	return c.EnqueueContext(context.Background(), task, opts...)
}
```

- `func (c *Client) EnqueueContext(ctx context.Context, task *Task, opts ...Option) (*TaskInfo, error)`

1. `composeOptions(opts ...Option) (option, error)` 将如入参多个 Option 组合成一个
2. 根据 Option 构建 `base.TaskMessage`
3. 根据任务执行时间和分组会有三种情况：
   - 情况一：以后执行，`client.schedule()`， state = TaskStateScheduled
   - 情况二：不是以后执行 && 配置了 groupOption，`client.addToGroup()`，state = TaskStateAggregating
     - 注意：会设置 `processAt = time.Time{}`，因为不知道合适聚合和处理任务
   - 情况三：立刻执行，`client.enqueue()`，state = TaskStatePending
     - 注意：会设置 `processAt = now`
4. `newTaskInfo()` 

## schedule

执行时间 > now ！

如果利用 `uniqueOption` 来保证任务的唯一性，也即利用 Redis 创建唯一锁来实现，会先计算当前任务的唯一锁的过期时间（`uniqueLockTTL = processAt.Add(uniqueTTL).Sub(now)`），然后调用 `ScheduleUnique()`； 否则调用 `Schedule()`

ScheduleUnique():
```shell
# 这里都是 Redis 命令，所有传入的动态值均以驼峰格式来表示

# 代码调用部分：
# 注册队列，将当前任务队列添加到 Redis 的 set 结构中，执行失败报错返回
SADD "asynq:queues" queueName

# Lua 脚本调用部分：
# 1. 将唯一 key 所代表的 taskID 存入 Redis 的 string 结构中
#    nx 代表 key 不存在时才设置，ex 指定过期时间
# 执行失败，代表 taskUniqueKey 已经存在，返回 -1
SET uniqueKey taskID nx ex uniqueLockTTL
# 2. 判断以下格式的 key 是否存在
# taskID 已经存在，返回 0
EXISTS "asynq:{queueName}:t:taskID"
# 3. 设置 taskID 所对应的 msg、state、unique_key，存入 Redis 的 hash 结构中
HSET "asynq:{queueName}:t:taskID" msg taskMessage state scheduled unique_key uniqueKey
# 4. 在 Redis 的 zset 结构中存入指定队列中 scheduled 状态下的 taskID，并设置分数（执行时间）
ZADD "asynq:{queueName}:scheduled" processAtSeconds taskID
# 5. 返回 1 
```

Schedule():
```shell
# 代码调用部分：
# 注册队列，将当前任务队列添加到 Redis 的 set 结构中，执行失败报错返回
SADD "asynq:queues" queueName

# Lua 脚本调用部分：
# 1. 判断以下格式的 key 是否存在
#  taskID 已经存在，返回 0
EXISTS "asynq:{queueName}:t:taskID"
# 2. 设置 taskID 所对应的 msg、state，存入 Redis 的 hash 结构中
# 相比于 ScheduleUnique，这里不存入 unique_key
HSET "asynq:{queueName}:t:taskID" msg taskMessage state scheduled
# 3. 在 Redis 的 zset 结构中存入指定 scheduled 状态的队列下的 taskID，并设置分数（执行时间）
ZADD "asynq:{queueName}:scheduled" processAtSeconds taskID
# 返回 1
```

## addToGroup

执行时间 <= now，且设置了 group！

如果利用 `uniqueOption` 来保证任务的唯一性，也即利用 Redis 创建唯一锁来实现，调用 `AddToGroupUnique()`，否则调用 `AddToGroup()`。

AddToGroupUnique():

注意，这里的 uniqueLockTTL 和 `ScheduleUnique()` 不同，因为执行时间 <= now，所以不再计算，而是使用配置的 ttl！

```shell
# 代码调用部分：
# 注册队列，将当前任务队列添加到 Redis 的 set 结构中，执行失败报错返回
SADD "asynq:queues" queueName

# Lua 脚本调用部分：
# 1. 将唯一 key 所代表的 taskID 存入 Redis 的 string 结构中
#    nx 代表 key 不存在时才设置，ex 指定过期时间
# 执行失败，代表 taskUniqueKey 已经存在，返回 -1
SET uniqueKey taskID nx ex uniqueLockTTL
# 2. 判断以下格式的 key 是否存在
#  taskID 已经存在，返回 0
EXISTS "asynq:{queueName}:t:taskID"
# 3. 设置 taskID 所对应的 msg、state、group，存入 Redis 的 hash 结构中
# 注意，这里没有放入 unique_key
HSET "asynq:{queueName}:t:taskID" msg taskMessage state aggregating group groupKey
# 4. 在 Redis 的 zset 结构中存入指定队列中指定 group 中的 taskID，并设置分数（执行时间）
ZADD "asynq:{queueName}:g:groupKey" currentUnixTime taskID
# 5.向同一队列下的 groups 中添加新的 group，这里是在记录同一队列下的所有 group
SADD "asynq:{queueName}:groups" groupKey
# 6. 返回 1 
```

AddToGroup():

```shell
# 代码调用部分：
# 注册队列，将当前任务队列添加到 Redis 的 set 结构中，执行失败报错返回
SADD "asynq:queues" queueName

# Lua 脚本调用部分：
# 1. 判断以下格式的 key 是否存在
# taskID 已经存在，返回 0
EXISTS "asynq:{queueName}:t:taskID"
# 2. 设置 taskID 所对应的 msg、state、group，存入 Redis 的 hash 结构中
HSET "asynq:{queueName}:t:taskID" msg taskMessage state aggregating group groupKey
# 3. 在 Redis 的 zset 结构中存入指定队列中指定 group 中的 taskID，并设置分数（执行时间）
ZADD "asynq:{queueName}:g:groupKey" currentUnixTime taskID
# 4.向同一队列下的 groups 中添加新的 group，这里是在记录同一队列下的所有 group
SADD "asynq:{queueName}:groups" groupKey
# 5. 返回 1 
```

## enqueue

执行时间 <= now，且未设置了 group！

如果利用 `uniqueOption` 来保证任务的唯一性，也即利用 Redis 创建唯一锁来实现，调用 `EnqueueUnique()`，否则调用 `Enqueue()`。

EnqueueUnique():

注意，这里的 uniqueLockTTL 和 `ScheduleUnique()` 不同，因为执行时间 <= now，所以不再计算，而是使用配置的 ttl！

```shell
# 代码调用部分：
# 注册队列，将当前任务队列添加到 Redis 的 set 结构中，执行失败报错返回
SADD "asynq:queues" queueName

# Lua 脚本调用部分：
# 1. 将唯一 key 所代表的 taskID 存入 Redis 的 string 结构中
#    nx 代表 key 不存在时才设置，ex 指定过期时间
# 执行失败，代表 taskUniqueKey 已经存在，返回 -1
SET uniqueKey taskID nx ex uniqueLockTTL
# 2. 判断以下格式的 key 是否存在
#  taskID 已经存在，返回 0
EXISTS "asynq:{queueName}:t:taskID"
# 3. 设置 taskID 所对应的 msg、state、pending_since、unique_key，存入 Redis 的 hash 结构中
HSET "asynq:{queueName}:t:taskID" msg taskMessage state pending pending_since currentUnixTimeNsec unique_key uniqueKey
# 4. 在 Redis 的 list 结构中，向左侧插入 taskID
LPUSH "asynq:{uniqueName}:pending" taskID
# 5. 返回 1 
```

Enqueue():

```shell
# 代码调用部分：
# 注册队列，将当前任务队列添加到 Redis 的 set 结构中，执行失败报错返回
SADD "asynq:queues" queueName

# Lua 脚本调用部分：
# 1. 判断以下格式的 key 是否存在
#  taskID 已经存在，返回 0
EXISTS "asynq:{queueName}:t:taskID"
# 2. 设置 taskID 所对应的 msg、state、pending_since，存入 Redis 的 hash 结构中
HSET "asynq:{queueName}:t:taskID" msg taskMessage state pending pending_since currentUnixTimeNsec
# 3. 在 Redis 的 list 结构中，向左侧插入 taskID
LPUSH "asynq:{uniqueName}:pending" taskID
# 4. 返回 1 
```

# Server

- `func NewServer(r RedisConnOpt, cfg Config) *Serve` 创建 Server，Server 中会包含
- Config 可以指定任务最大并发数、重试的延迟函数（计算重试间隔）等。

```go
// Server 负责任务处理和任务生命周期管理。
// 从 queue 中 pull task 并处理，如果某个任务处理未成功，将会调度以重试
type Server struct {
	logger *log.Logger
	// Redis client，下面那些功能都会持有这同一个 Redis Client
	broker base.Broker
	// 记录 Server 状态，可选值为 New、Active、Stopped、Closed，
	// Active 代表 server 已经在运行中了
	// Stopped 代表 server 已收到 SIGTSTP 信号，正在等待 shutdown
	// server.Run() 时会检查该状态，如果不是 New 会返回错误，如果是 New，更新为 Active
	state *serverState

	// 用于控制下面那些功能各自的 goroutine
	wg            sync.WaitGroup
	// 这些功能都会接受 Server Shutdown 时的信号，以便结束 goroutine
	forwarder     *forwarder // 定期检查所有 queue 中 scheduled 和 retry 集合，将准备好执行的任务移动到 pending 集合中
	processor     *processor // 根据优先级获取排序后的 queueNames，
	syncer        *syncer // 定期丢弃掉旧请求
	heartbeater   *heartbeater // 检查每个 worker 的租约是否有效，并定期心跳检查
	subscriber    *subscriber // 订阅 Redis 的 "asynq:cancel" 频道（channel），如果该 channel 有接收到消息（taskId），拿到该 id 对应的 cancel 函数，并执行，这样对应的任务会由于 context 被 cancel 而结束 
	recoverer     *recoverer // 定期获取所有队列中 30s 前已过期的 task，如果重试次数达到最大值，存档；否则重试
	healthchecker *healthchecker // 健康检查，定期去 ping Redis
	janitor       *janitor // 定期删除队列中 completed 集合中的过期任务
	aggregator    *aggregator // 将同一个 group 下的 task 聚合？
}
```

# 命令行

```shell
asynq stats   # 查看队列状态
```
