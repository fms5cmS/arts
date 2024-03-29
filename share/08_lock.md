
# 互斥vs可重入

互斥锁：保证只有一个进程或线程在临界区内。当锁被一个线程获取时，其他只能一直等待。

可重入锁(Reentrant Lock)：当一个线程获取锁时，如果没有其它线程拥有这个锁，那么，这个线程就成功获取到这个锁。之后，如果其它线程再请求这个锁，就会处于阻塞等待的状态。但是，如果拥有这把锁的线程再请求这把锁的话，不会阻塞，而是成功返回。

可重入锁解决了代码重入或者递归调用带来的死锁问题。如果持有可重入锁的话，可以可着劲儿地调用，比如通过递归实现一些算法，调用者不会阻塞或者死锁。

当持有锁的线程或 goroutine 释放锁，等待中的线程或 goroutines 中哪一个优先获取锁？常见互斥锁的实现：

- Barging：当锁被释放时，它会唤醒第一个等待者，然后把锁给第一个等待者或者给第一个请求锁的人。这种模式是为了提高吞吐量。
- HandsOff：当锁释放时候，锁会一直持有直到第一个等待者准备好获取锁。
  - 一个互斥锁的 HandsOff 会完美地平衡两个 goroutines 间的锁分配，但是会降低性能，因为它会迫使第一个 goroutine 等待锁。
- Spinning：自旋在等待队列为空或者应用程序重度使用锁时效果不错。
  - parking 和 unparking goroutines 有不低的性能成本开销，相比自旋来说要慢得多。

> goroutine park：runtime.gopark 将当前 goroutine 挂起并解绑 G 和 M，重新调度其他的 G

# 死锁

死锁必须同时满足的四个条件：

- 互斥： 至少一个资源是被排他性独享的，其他线程必须处于等待状态，直到资源被释放。
- 持有和等待：线程持有一个资源，并且还在请求其它线程持有的资源
- 不可剥夺：资源只能由持有它的线程来释放。
- 环路等待：一般来说，存在一组等待进程，P={P1，P2，…，PN}，P1 等待 P2 持有的资源，P2 等待 P3 持有的资源，依此类推，最后是 PN 等待 P1 持有的资源，这就形成了一个环路等待的死结

# 分布式锁

[Redis 分布式锁](./01_distributedLock.md)

# MySQL 中的锁

根据加锁的范围，MySQL 里面的锁大致可以分成全局锁、表级锁和行锁三类。

全局锁是对整个数据库实例加锁，如在做全库逻辑备份时会使用。

## 表锁

MySQL 中有两种表级锁：表锁、元数据锁(Meta Data Lock, MDL)。

表锁的语法：`LOCK TABLES ... READ/WRITE`，除了会限制别的线程读写外，也先定了当前线程接下来的操作对象。

```sql
-- 其他线程写 table1、读写 table2 的语句都会被阻塞
-- 当前线程执行 UNLOCK TABLES 之前，页只能执行度 table1、读写 table2，不允许写 table1
LOCK TABLES table1 READ, table2 WRITE;
```

一般也不使用 `LOCK TABLES` 来控制并发，锁住整个表的影响太大。

---

MySQL5.5 中引入了 MDL，MDL 不需要显式使用，在访问一个表时会被自动加上，用于保证读写的正确性。

当对一个表做增删改查时，加 MDL 读锁；做表结构变更操作时，加 MDL 写锁。

"给表加字段导致长时间阻塞"示例：

| sessionA                       | sessionB                       | sessionC                       | sessionD                       |
|--------------------------------|--------------------------------|--------------------------------|--------------------------------|
| `BEGIN;`                       |                                |                                |                                |
| `SELECT * FROM table LIMIT 1;` |                                |                                |                                |
|                                | `SELECT * FROM table LIMIT 1;` |                                |                                |
|                                |                                | `ALTER TABLE table ADD f INT;` |                                |
|                                |                                |                                | `SELECT * FROM table LIMIT 1;` |

sessionA 先启动，对 table 表加一个 MDL 读锁，而 sessionB 由于需要的也是 MDL 读锁，因此可以正常执行；

由于 sessionA 的 MDL 读锁未被释放，而 sessionC 需要 MDL 写锁，所以sessionC 也被阻塞住；

之后的 sessionD 需要申请 MDL 读锁，也会被 sessionC 阻塞！如果该表查询语句频繁，且客户端有重试机制（超时后会再起一个 session 再请求），这个库的线程很快就会用完。

> 事务中的 MDL 锁，在语句执行开始时申请，但是语句结束后不会马上释放，而是等到事务提交后再提交。
> 
> 在 MySQL 的 information_schema 库的 innodb_trx 表中，可以查到当前执行中的事务。
> 
> 如果要做 DDL 变更的表刚好有长事务在执行，需要考虑先暂停 DDL 或 kill 掉这个长事务

那么，如何安全地给小表加字段？

首先，需要解决长事务，事务不提交，就会一直占着 MDL 锁。

如果变更的是一个热点表，虽然数据量不大，但请求很频繁，而又必须要加一个字段，这个时候 kill 可能未必管用，因为新的请求很快就来了。

比较理想的机制：在 `ALERT TABLE` 中设定等待时间，如果在指定的等待时间内可以拿到 MDL 写锁最好，拿不到也不要阻塞后面的业务语句，先放弃。之后 DBA 再通过重试命令重复这个过程。

> MariaDB 支持 DDL NOWAIT/WAIT n 的语法

---

MySQL 插入数据时，自增列的值可以不指定，MySQL 会自动填充。而这个自增值的获取有两种方式（innodb_autoinc_lock_mode 来配置）：

- innodb_autoinc_lock_mode=0，利用表锁。在执行插入语句时会加一个表级锁，然后为每条插入记录的自增列分配自增的值，当语句结束后再释放锁
  - 注意：并不是等到事务结束后才释放
- innodb_autoinc_lock_mode=2，使用轻量级锁，避免表锁。在为插入语句生成自增列的值时获取这个锁，然后在生成值后就会释放锁，不需要等到整个插入语句执行完成后才释放锁
- innodb_autoinc_lock_mode=1，以上方式混用，插入记录的数量确定时使用轻量级锁，否则使用表锁

> MySQL5.7 及之前版本，自增值保存在内存中，没有持久化。
> 
> 如果一个表当前数据行里最大的 id 是 10，内存中记录的 AUTO_INCREMENT=11。这时删除 id=10 的行，AUTO_INCREMENT 还是 11。但如果马上重启实例，重启后这个表的 AUTO_INCREMENT 就会变成 10。
> 
> MySQL8.0 起，将自增值的变更记录在了 redo log 中，重启的时候依靠 redo log 恢复重启之前的值

唯一键冲突、事务回滚会导致自增 id 不连续：

`INSERT INTO table(id, v1, v2) VALUES(NULL, 1, 1);` 该语句的流程（使用的是轻量级锁）：

1. 执行器调用 InnoDB 引擎接口写入一行，传入的这一行的值是 (0,1,1);
2. InnoDB 发现用户没有指定自增 id 的值，获取表 t 当前的自增值 2；
3. 将传入的行的值改成 (2,1,1);
4. 将表的自增值改成 3；
5. 继续执行插入数据操作，由于已经存在 c=1 的记录，所以报 Duplicate key error，语句返回。

而上面步骤 4 并不会在语句执行失败后把自增值改回去，所以之后再次插入时，自增值就会跳过 2。

## 行锁

MySQL 的行锁是在引擎层由各个引擎自己实现的。如 MyISAM 就不支持行锁。

在 InnoDB 事务中，行锁是在需要的时候才加上的，但并不是不需要了就立刻释放，而是要**等到事务结束时才释放**。这个就是两阶段锁协议。所以，**如果事务中需要锁多个行，要把最可能造成锁冲突、最可能影响并发度的锁尽量往后放**，减少这一行的锁时间，从而减少事务间的锁等待。

InnoDB 中的行锁：

- LOCK_REC_NOT_GAP：仅对记录本身加锁
- LOCK_GAP：锁住记录前的间隙，防止别的事务向这个间隙插入新记录
  - 如对 id 在 (3, 8) 区间内加锁后，在锁释放前，不允许向这个范围内的插入记录
  - 仅仅是为了防止插入幻读的记录
- LOCK_ORDINARY：相当于上面两种锁的结合，既可以保护记录本身，也可以防止别的事务向间隙插入新记录
  - 如对 id 在 (3, 8] 区间内加锁 
- LOCK_INSERT_INTENTION（插入意向锁）：在当前事务插入记录时因碰到别的事务加的 gap 锁而进入等待状态时所生成的一个锁结构
- 隐式锁：依靠记录的 trx_id 来保护不被别的事务改动该记录

## 死锁&死锁检测

MySQL 遇到死锁后，有两种策略：

- 直接进入等待，直到超时。这个超时时间可以通过参数 innodb_lock_wait_timeout 来设置（默认 50s）
- 发起死锁检测，发现死锁后，主动回滚死锁链条中的某一个事务，让其他事务得以继续执行
  - 默认参数 innodb_deadlock_detect 为 on，表示开启这个逻辑。

> 一致性读不会加锁，就不需要做死锁检测

正常情况下都是要采用第二种策略的，因为设置超时时间过长对服务影响不可接受，过短又容易误判。但是，每当一个事务被锁时，都需要判断它所在的线程有没有被别人锁住。所以死锁检测会耗费大量的 CPU 资源！

所以可以通过控制并发度，这样死锁检测的成本也会降低，而这个并发控制需要在数据库层服务端来做，如通过中间件或修改 MySQL 源码，来将对同行的更新，在进入引擎之前排队。这样在 InnoDB 内部就不会有大量的死锁检测工作了。

可以通过 `SHOW ENGINE INNODB STATUS` 语句查看最近发生的一次死锁信息。尽可能避免死锁

# Golang 中的锁

## Mutex

Golang 中的 `sync.Mutex` 结构：

```go
package sync

type Mutex struct {
   // 其二进制形式的最低几位：
   // 第一位（最低位）代表了该锁是否被 goroutine 持有，即是否已上锁；
   // 第二位代表了是否有等待的 goroutine 被唤醒；
   // 第三位代表了该锁是否处于饥饿模式；
   // 其余位数所对应的十进制数代表了等待该锁的 goroutine 数。
   // 假设 state==23，其对应的二进制数位 10111，去掉前三位后，剩余的二进制数为 10，对应的十进制数是 2，所以等待该锁的 goroutine 共有 2 个。
   state int32
   sema  uint32
}

const (
   mutexLocked = 1 << iota // mutex is locked
   mutexWoken
   mutexStarving
   mutexWaiterShift = iota
   starvationThresholdNs = 1e6
)

```

### 发展

Go1.8 的互斥锁使用了 Barging 和 Spinning 的结合实现。当试图获取已经被持有的锁时，如果本地队列为空并且 P 的数量大于 1，goroutine 将自旋几次（用一个 P 自旋会阻塞程序）。自旋后，goroutine park。

饥饿问题：由于新来的 goroutine 也参与竞争，有可能每次都被新来的 goroutine 抢到锁，极端情况下，等待中的 goroutine 可能会一直获取不到锁。

> 新来的 goroutine 有先天优势，因为它们正在 CPU 中运行。

Go1.9 新增了饥饿模式解决上述问题，这种模式会在锁释放时触发 HandOff。

所有等待锁时长超过 1ms 的 goroutine（也称为有界等待）将被诊断为饥饿。当被标记为饥饿状态时，`Unlock()` 方法会 HandOff 把锁直接扔给第一个等待者。

Go1.18 中给原生的 Mutex 添加了 `TryLock()` 的功能，用于非阻塞式的获取锁，如果没有获取到锁，会直接返回 false，而不会一直阻塞。

### 正常模式&饥饿模式

- 正常模式 -> 饥饿模式：

正常模式，waiter 都是进入先进先出队列，被唤醒的 waiter 不会直接持有锁，而要和新来的 goroutine 竞争。高并发情况下，被唤醒的 waiter 可能获取不到锁，这时，它会被插入到队列前面，如果 waiter 获取不到锁的时间超过阈值 1ms，则当前 Mutex 就进入到了饥饿模式。

在饥饿模式下，Mutex 的拥有者将直接把锁交给队列最前面的 waiter。新来的 goroutine 不会尝试获取锁，即使看起来锁没有被持有，它也不会去抢，也不会 spin，它会乖乖地加入到等待队列的尾部。

- 饥饿模式 -> 正常模式：

1. 此 waiter 已经是队列中的最后一个 waiter 了，没有其它的等待锁的 goroutine 了
2. 此 waiter 的等待时间小于 1ms

满足以上任意一个条件，Mutex 就会从饥饿模式转为正常模式。

饥饿模式是对公平性和性能的一种平衡，它避免了某些 goroutine 长时间的等待锁。在饥饿模式下，优先对待的是那些一直在等待的 waiter。

### 使用

[Mutex 使用示例](./examples/lock/mutex_test.go)

Go 在运行时有死锁检查机制，能够发现死锁问题。

为了及时发现问题，可以使用 vet 工具，`go vet xx.go` 来检查，并将其放入 Makefile 文件中，在持续集成的时候跑一跑，这样可以及时发现问题，及时修复。

- 易错场景一：`Lock()`、`Unlock()` 没有成对出现

可能会导致死锁，又或者由于 `Unlock()` 一个未知的锁导致 panic。

- 易错场景二：Copy 一个已使用的锁

Mutex 是一个有状态的对象，它的 state 字段记录这个锁的状态。如果要复制一个已经加锁的 Mutex 给一个新的变量，那么新的刚初始化的变量就已经被加锁了！

这种场景下也容易导致死锁，可以提前通过 Go 的 vet 工具来检测是否进行了 Mutex 复制的问题。

> sync 包下的同步原语都是不允许复制的！

- 易错场景三：重入

Mutex 是[互斥锁，而非可重入锁](#互斥vs可重入)

Golang 的 Mutex 并没有标识锁当前是被哪个 goroutine 持有的，所以也就不具备可重入的能力！

如果进行了 Mutex 的重入，就会发现程序一直尝试获取锁但是又获取不到，导致发生死锁。见[代码示例](./examples/lock/reentrantLock_test.go)

如何实现可重入锁？其关键是记录具体哪个 goroutine 当前持有锁。

- 易错场景四：死锁

## RWMutex

RWMutex 一般都是基于互斥锁、条件变量（condition variables）或者信号量（semaphores）等并发原语来实现。Go 标准库中的 RWMutex 是基于 Mutex 实现的。

Go 标准库中的 RWMutex 设计是 Write-preferring 方案。一个正在阻塞的 Lock 调用会排除新的 reader 请求到锁。

> 读写锁的设计和实现分为三种：
> 
> Read-preferring，读优先：可以提供很高的并发性，但是，在竞争激烈的情况下可能会导致写饥饿。
> 
> Write-preferring，写优先：如果已经有一个 writer 在等待请求锁的话，它会阻止新来的请求锁的 reader 获取到锁，所以优先保障 writer。当然，如果有一些 reader 已经请求了锁的话，新请求的 writer 也会等待已经存在的 reader 都释放锁之后才能获取。
> 
> 不指定优先级：这种设计比较简单，不区分 reader 和 writer 优先级

```go
type RWMutex struct {
  w           Mutex   // 互斥锁解决多个 writer 的竞争
  writerSem   uint32  // writer 信号量
  readerSem   uint32  // reader 信号量
  readerCount int32   // reader 的数量（以及是否有 writer 竞争锁）
  readerWait  int32   // writer 请求锁时需要等待完成 read 的 reader 的数量
}
```

如果遇到可以明确区分 reader 和 writer goroutine 的场景，且有大量的并发读、少量的并发写，并且有强烈的性能需求，你就可以考虑使用读写锁 RWMutex 替换 Mutex。

[使用示例](./examples/lock/rwmutex_test.go)

## channel 实现锁

channel 实现锁有两种方式：

- 初始化一个 capacity=1 的 channel，并向其中放入一个元素，该元素就代表了锁，谁取得这个元素就相当于获得了锁，[示例代码](./examples/lock/channelMutex.go)
- 初始化一个 capacity=1 的 channel，它的空槽就代表了锁，谁能成功向其发送元素，谁就获得了锁

## 竞态检查

Go 提供了一个检测并发访问共享资源是否有问题的工具 race detector，可以帮助我们自动发现程序有没有 data race 的问题。

编译器通过探测所有的内存访问，加入代码能监视对这些内存地址的访问（读还是写）。在代码运行的时候，race detector 就能监控到对共享变量的非同步访问，出现 race 的时候，就会打印出警告信息。

使用：在使用命令行来编译(compile)、测试(test)、运行(run) 时，加上 `-race` 参数即可。

由于其实现方式，只能通过真正对实际地址进行读写访问的时候才能探测，所以它并不能在编译的时候发现 data race 的问题。而且，在运行的时候，只有在触发了 data race 之后，才能检测到，如果碰巧没有触发的话，是检测不出来的。

  > 通过在编译的时候插入一些指令，在运行时通过这些插入的指令检测并发读写从而发现 data race 问题，就是这个工具的实现机制
