
Go1.21.5 内置库 sync

注意：**内置的 sync 包中的同步原语在使用后是不能复制的**！！！如果需要作为参数传递的话请使用指针。

Go 内置的锁都是[互斥锁，不是可重入锁](../../share/08_lock.md#互斥vs可重入)！

# sync.Mutex

见[Mutex](../../share/08_lock.md#mutex)

# sync.RWMutex

见[RWMutex](../../share/08_lock.md#rwmutex)

# sync.WaitGroup

WaitGroup 是用来做任务编排的。问题：goroutine 要进行任务 A，必须等待其他 goroutines 完成任务 B、C、D 等后才可以开始。

注意：如果在 goroutine 中调用 `Add()` 方法，可能会导致子 goroutine 没有执行，就已经完成了主任务。

[代码示例](./waitgroup_test.go) 中的 TestWaitGroup_Err，由于四个 goroutine 一开始都休眠，所以可能在调用 `Wait()` 时，四个 `Add()` 还没有调用，所以并没有等待四个子任务执行就立刻执行了下一步。

# sync.Once

once.go：

```go
type Once struct {
   done uint32 // 只会有 0 和 1 两种值
   m    Mutex
}

func (o *Once) Do(f func()) {
   // 不能使用 CAS 操作！ 
   // if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
   //    f()
   // }
   // 这是为了保证 Do() 函数返回时 f 一定执行完成了！
   // 说明：
   // 假设有 A、B 两个 goroutine 同时读到了 done 字段为 0
   // 而 A 先成功完成了 CAS 操作将 done 置一，注意：此时 f 函数尚未执行完成
   // 然后 B 开始进行 CAS 操作，由于 done 的值变成了 1，与之前的 0 不等，操作失败 return
   // B return 后认为 f 函数已经执行完成了，然后访问 f 初始化的全局唯一资源，如数据库连接池
   // 程序就会 panic，因为 done 的置一操作是在 f 执行完之前就已经完成了
   if atomic.LoadUint32(&o.done) == 0 {
      // Outlined slow-path to allow inlining of the fast-path.
      o.doSlow(f)
   }
}
// 这部分操作单独拆分出来了！让 Do 函数可以被内联到具体的调用处
func (o *Once) doSlow(f func()) {
   o.m.Lock()
   defer o.m.Unlock()
   if o.done == 0 {
      // 注意这里的顺序，先执行 f 函数，然后再将 done 置一的
      defer atomic.StoreUint32(&o.done, 1)
      f()
   }
}
```

> 如果有多个只需执行一次的函数，就应该为每个函数都分配一个 sync.Once
>
> 包含以下内容的方法都不会被内联：闭包调用，select，for，defer，go关键字创建的协程

假设在 `Call()` 中进行了 `once.Do()` 的调用:

- doSlow 不拆出来而是把逻辑放在了 Do 里面： 
  - 会因为里面有 defer 操作而导致整个 Do 函数没办法被内联，所以无论是什么时候调用 f，都必然会有一次 Do 的函数调用 
  - 无论是读到 done==0 还是 done==1 的 goroutine 的函数调用次数都是 1
- doSlow 拆分出来：
  - Do 会被内联到具体的调用处，这样的话 doSlow 只有读到 done==0 的 goroutine 才会额外有 doSlow 的函数调用，一旦 done 置为 1 后，再调用的话就不会再有 doSlow 的函数调用了
  - 只有读到 done==0 goroutine 的函数调用次数是 1，读到 done==1 的goroutine 的函数调用次数是 0

Go 1.21.0 中增加了三个 Once 相关的函数（不是 Once 的方法，位于 oncefunc.go 文件中）对 Once 进行了封装用于简化 sync.Once 的调用：

- `OnceFunc(f func()) func()`：返回函数 g，多次调用 g 只会执行一次 f
  - 如果 f 执行时 panic, 则后续调用函数 g 不会再执行 f,但是每次调用都会 panic
- `OnceValue[T any](f func() T) func() T`：同上，且调用 g 返回的值为 f 执行后返回的值
- `OnceValues[T1, T2 any](f func() (T1, T2)) func() (T1, T2)`：同上，多了一个返回值

# sync.Pool

sync.Pool 用于保存一组可独立访问的**临时**对象，可以创建池化的对象，但是，它池化的对象可能会无通知地被 GC 回收掉，这对于数据库等长连接场景是不合适的！！

> sync.Pool 本身是线程安全的，多个 goroutines 可以并发地调用它的方法存取对象；
> 
> sync.Pool 不可在使用后再复制使用。

Pool 仅提供了三个对外的方法：

- `New` 该方法包含在 Pool 结构体中作为一个字段，当调用 `Get()` 从池中获取元素但没有空闲元素返回时会调用 `New()` 来创建新的元素。
  - 建议在创建 Pool 对象时设置该字段的值
- `Get()` 从池中取走一个元素，当没有设置 New 字段或池中没有空闲元素时，返回值为 nil，因此使用时可能需要进行判断
- `Put()` 将元素放入池中，如果放入一个 nil，则 Pool 会忽略这个值

保存和复用临时对象，减少内存分配，主要使用场景：
- 进程中的 inuse_objects 数过多，内存占用越来越高（监控图上进程 RSS 占用过高）
- 进程中的 inuse_objects 数过多，gc mark（需要遍历所有存活的对象） 消耗大量 CPU

如 buffer 池，[示例](./pool_test.go)

## 实现

```go
type Pool struct {
    noCopy noCopy
    // 指向了一个数组，该数组大小固定，长度为 P 的数量（GOMAXPROCS）
    // 实际类型为 [P的数量]poolLocal
    // 当前运行 P 要从 Pool 中获取对象，需要从指定的索引下获取
    local     unsafe.Pointer
    localSize uintptr        // size of the local array
    // 发生 GC 时，会将 local 赋值给 victim，localSize 赋值给 victimSize，如果原本有值，会将原本值丢弃
    // 获取对象时，如果 local 为空，那么就会从 victim 中找，逻辑和从 local 获取一致
    victim     unsafe.Pointer // local from previous cycle
    victimSize uintptr        // size of victims array
    New func() interface{}
}

// 本地 P 作为唯一的生产者，多个 P 作为消费者
type poolLocalInternal struct {
    private any       // 仅在本地 P 上使用，无锁操作，没有并发问题
    shared  poolChain // 无锁队列， Local P can pushHead/popHead; any P can popTail.
}

type poolLocal struct {
    poolLocalInternal
    // CPU 缓存对其，从而避免 false sharing 
    pad [128 - unsafe.Sizeof(poolLocalInternal{})%128]byte
}
```

每次创建 Pool 对象时，会将其追加到全局的 runtime.allPools 对象中！

Pool 最重要的两个字段是 local、victim。

每次 GC 时，Pool 会把 victim 中的对象移除，然后把 local 的数据给 victim。victim 中的元素可能会被当作垃圾丢弃，也可能被捡回来重新使用。

> `pin()` 方法会将当前 goroutine 固定在当前 P 上，避免查找元素期间被其他的 P 执行。

- Get

1. 先从本地的 private 字段中获取可用元素
2. 如果没有获取到，尝试从 shared 中获取
3. 还是没有获取到，使用 `getSlow()` 从其他 P 的 shared 中偷一个
   1. 遍历所有的 local，尝试从它们的 shared 弹出其末尾的一个元素
   2. 如果一个都没找到，开始查询 victim，查询逻辑一样，先从对应的 victim 的 private 查找，没有的话从其他 victim 的 shared 中查找
4. 依然没有获取到，尝试使用 New 创建一个新的元素

- Put

1. 设置本地 private 字段
2. private 字段有值，将元素加入到本地 shared 队列的头部

## 踩坑

- 内存泄漏

取出来的 `bytes.Buffer` 在使用的时候，我们可以往这个元素中增加大量的 byte 数据，这会导致底层的 byte slice 的容量可能会变得很大。

这个时候，即使 Reset 再放回到池子中，这些 byte slice 的容量不会改变，所占的空间依然很大。而且，因为 Pool 回收的机制，这些大的 Buffer 可能不被回收，而是会一直占用很大的空间，这属于内存泄漏的问题。

解决方法：在回收 buffer 时，一定要检查回收对象的大小，如果 buffer 太大，就不要回收了。

- 内存浪费

池子中的 buffer 都比较大，但实际使用时，很多时候仅需一个小的 buffer，这就造成了浪费。

解决方法：可以将 buffer 池分为几层，如小于 512 byte 的元素 buffer 一个池子；小于 1k byte 的元素一个池子；小于 4k byte 的元素一个池子，按需从对应池子中获取 buffer

# sync.Map

见[sync.Map](../../share/06_hash.md#syncmap)

# sync.Cond（很少用，很难掌握）

Cond 是为等待/通知场景下的问题提供支持的。问题：满足某个触发条件时，通知等待的多个 goroutine 去执行。

很少使用，一般遇到需要使用 Cond 的场景，更多会使用 channel 的方式。且很难掌握，[代码示例](./cond_test.go)。

# sync/atomic 包

sync/atomic 包提供了一些实现原子操作的方法。

使用场景示例：使用一个标志来标识一个定时任务是否已经启动执行。

可以使用加锁的方式，在读取、设置标志时加锁，保证同一时刻只有一个定时任务在执行。

由于该场景中不涉及对资源复杂的竞争逻辑，只会并发地读写标志，这类场景就适合使用 atomic 的原子操作：可以使用一个 uint32 的变量，如果该变量值为 0 则代表没有任务在执行，为 1 则标识已有任务在执行了。

此外 atomic 原子操作还是实现 lock-free 数据结构的基础！[无锁队列](lkQueue.go)

常用于 COW（Copy On Writer）：把老的对象复制一份，添加新的数据进去，再利用原子操作（`atomic.Value` 类型）将老的对象替换掉。

见 [COW 示例 TestCOW](atomic_test.go)

> COW 的使用：
>
> Redis 的 `bgsave` 命令就是 fork 一个子进程来进行持久化（RDB 方式），子进程先将数据写入一个临时文件中，待持久化过程都结束了，再用这个临时文件替换上次持久化好的文件。整个过程中，主进程是不进行任何 IO 操作的，这就确保了极高的性能。
> 
> -----> 扩展：[Redis 的RDB持久化方式](../../share/05_persistence.md#rdb)
> 
> Docker 的 UnionFS 也使用了 COW

注意：atomic 操作的对象是一个地址，需要把可以寻址的变量的地址作为参数传递给方法，而不是把变量的值传递给方法。

方法：

- `AddXX(addr, delta)` 给第一个参数地址中的值增加一个 delta 值，返回计算结果
  - 对于有符号整数而言，delta 可以是负数，相当于做减法
  - 对于无符号整数，如何做减法呢？见 [TestAtomicSub](atomic_test.go)
- `CompareAndSwapXX(addr, old, new)` CAS 操作，返回操作是否成功
- `SwapXX(addr, new)` 进行替换操作，返回旧值
- `LoadXX(addr)` 取出给定地址中的值
- `StoreXX(addr, val)` 把值存入指定地址中

类型：

- Value，可以原子地存取对象类型，仅存取，常用语配置变更等场景，见 [TestAtomicValue](atomic_test.go)
- Uintptr
- Uint64
- Uint32
- Int64
- Int32
- Pointer
- Bool
