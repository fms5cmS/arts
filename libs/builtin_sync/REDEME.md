
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


# sync.Cond（很少用，很难掌握）

Cond 是为等待/通知场景下的问题提供支持的。问题：满足某个触发条件时，通知等待的多个 goroutine 去执行。

很少使用，一般遇到需要使用 Cond 的场景，更多会使用 channel 的方式。且很难掌握，[代码示例](./cond_test.go)。

