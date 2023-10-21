
传统的单体架构的时候，基本是单库然后业务单表的结构。每个业务表的ID一般都是从 1 增，通过 `AUTO_INCREMENT=1` 设置自增起始值。而在分布式服务架构模式下分库分表的设计，使得多个库或多个表存储相同的业务数据。这种情况根据数据库的自增ID就会产生相同ID的情况，不能保证主键的唯一性。

# 数据库自增ID

```
DB1 从 1 开始自增，每次自增 3：1, 4, 7, 10...

DB2 从 2 开始自增，每次自增 3：2, 5, 8, 11...

DB3 从 3 开始自增，每次自增 3：3, 6, 9, 12...
```

如 MySQL：

- `auto_increment_increment`，设置自增字段的起始值
- `auto_increment_offset`，设置每次自增的量

该方式强依赖 DB，数据库主从切换时的不一致可能导致重复发号。

# Redis

Redis 提供了类似 `INCR`、`INCRBY` 的自增原子命令，且由于 Redis 自身的单线程特点，可以保证生成的 ID 是唯一有序的。

单机存在性能瓶颈，集群方式会设计到和数据库集群同样的问题，所以也需要设置分段和步长来实现。

为了避免长期自增后数字过大可以通过与当前时间戳组合起来使用，另外为了保证并发和业务多线程的问题可以采用 Redis + Lua的方式进行编码，保证安全。

Redis 实现分布式全局唯一ID，它的性能比较高，生成的数据是有序的，对排序业务有利，但是同样它依赖于 Redis，需要系统引进 Redis 组件，增加了系统的配置复杂性。

# uuid

[uuid](https://github.com/google/uuid)(Universally Unique Identifier，通用唯一识别码) 是一个 32 长度的字符串（每个字符 16 进制，所以是 128bits），格式为 8-4-4-4-12，如 `06722a93-7b8e-46f6-ade4-88e2e15ffbff`，除连字符 `-` 外的每个字符都是 16 进制数字，理论上最大长度为 16^32。一般使用时会删除 `-`，如 `strings.ReplaceAll(uuid.NewString(), "-", "")`。

uuid 的产生方式有五个版本：

- v1，基于时间：一般通过当前时间、随机数、本地 MAC 地址计算，由于使用了 MAC 地址，因此可以确保唯一性，但也暴露了 MAC 地址，私密性不够好
- v2，DCE 安全：DCE(Distributed Computing Environment) 安全的 uuid 和基于时间的 uuid 算法相同，但会把时间戳的前 4 位置换为 POSIX 的 uid 或 gid，使用较少
- v3，基于名字（MD5）：通过计算名字和名字空间的 MD5 散列值得到
- v4，随机：根据随机数或伪随机数生成，重复的可能性可以忽略不计，经常使用
- v5，基于名字（SHA1）：和上一种类似，但散列值计算使用了 SHA1 算法

缺点：不易存储、信息不安全（指基于 MAC 地址生成时）、对 MySQL 索引不利。

# ulid

uuid v1/v2 在多数情况下是不实际的，如要求访问一个唯一稳定的 MAC 地址；uuid v3/v5 需要唯一种子并生成随机分布的 ID，这可能会导致很多数据结构的碎片；uuid v4 仅需提供随机性可能会导致数据结构碎片。

[ulid](https://github.com/oklog/ulid)(Universally Unique Lexicographically Sortable Identifier) 是一个规范编码为 26 个字符的字符串（不区分大小写，不包含特殊字符），每毫秒可以生成 1.21e+24 个唯一的 ulid，且是字典有序的，当正确检测并处理同一毫秒，ulid 是单调有序的。

ulid 使用了 [Crockford 的 base32](http://www.crockford.com/wrmg/base32.html) 提高效率和可读性（每个字符 5bits）。

相比 uuid v4 仅提供随机性，ulid 根据毫秒级的时间戳以及随机数构建。

ulid 是自动单调的，但仅限于毫秒的精度，如当下毫秒生成的所有 ulids 都大于前一毫秒内的所有 ulids。

而同一毫秒内生成的 ulids 按其随机量排序，这就意味着它们默认是无序的。可以使用 `ulid.MonotonicEntropy` 或 `ulid.LockedMonotonicEntropy` 来生成给定毫秒内单调的 ulids。

> go 的 ulid 实现 https://github.com/oklog/ulid 中时间戳为 48bits，可以使用到 10889 年，占前 10 个字符；熵为 80bits，占了后 16 个字符。
> 
> Crockford 的 base32 字母表仅使用 0123456789ABCDEFGHJKMNPQRSTVWXYZ，不使用 I、L、O、U 以避免混淆和滥用。

# snowflake

Twitter 的 snowflake 算法中，id 类型为 int64，被分为四部分（不含第一个 bit，该 bit 为符号位，值始终为 0，保证生成的 ID 为正数）：

- `timestamp`，41 位标识收到请求时的时间戳，时间单位为 ms，由程序在运行期生成
- `datacenter_id`，5 位标识数据中心或机房的 id
- `worker_id`，5 位标识机器的实例 id
- `sequence_id`，12 位的循环自增 id，达到 1111_1111_1111 后会归零，由程序在运行期生成

> 标识 timestamp 的 41 位可以支持使用 69 年，不必从 1970 年开始，可以自己指定起始时间；
>
> 数据中心加上实例 id 共计 10 位，可以支持每个数据中心部署 2^5=32 台机器，所有数据中心共 1024 台实例；
> 
> 同一台机器同一毫秒可以产生 2^12=4096 条消息，一秒内共计 409.6w 条数据。

四个字段中 `worker_id` 是逻辑上给机器分配的 id，该如何分配？可以由 MySQL 这种提供自增 id 功能的工具支持，且需要持久化，以避免每次上线时都需要获取新的 `worker_id`；也可以简单一些，把 `worker_id` 直接写到 worker 的配置中。

[snowflake](https://github.com/bwmarrin/snowflake) 是一个轻量化的 snowflake 的 Go 实现。

snowflake 不依赖数据库等第三方系统，以服务的方式部署，稳定性更高，生成 ID 的性能也很高，且可以根据自身业务特性分配 bit 位，非常灵活。

但是 snowflake 强依赖机器时钟，如果机器时钟回拨会导致发号重复或服务处于不可用状态！

# leaf

[leaf](https://github.com/Meituan-Dianping/Leaf) 是美团推出的分布式 ID 生成算法，提供了两种 ID 生成方案：

## leaf-segment

在[数据库方案](#数据库自增ID)中每次获取ID都得读写一次数据库，造成数据库压力大。改为利用 proxy server 批量获取，每次获取一个 segment(step决定大小) 号段的值。用完之后再去数据库获取新的号段，可以大大的减轻数据库的压力；

各个业务不同的发号需求用 biz_tag 字段来区分，每个 biz_tag 的 ID 获取相互隔离，互不影响。如果以后有性能需求需要对数据库扩容，不需要上述描述的复杂的扩容操作，只需要对 biz_tag 分库分表就行。

```sql
CREATE TABLE `leaf_alloc` (
  `biz_tag` varchar(128)  NOT NULL DEFAULT '' COMMENT '业务key',
  `max_id` bigint(20) NOT NULL DEFAULT '1' COMMENT '当前已经分配了的最大id',
  `step` int(11) NOT NULL COMMENT '初始步长，也是动态调整的最小步长',
  `description` varchar(256)  DEFAULT NULL COMMENT '业务key的描述',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`biz_tag`)
) ENGINE=InnoDB;
```

为了解决 TP999（满足千分之九百九十九的网络请求所需要的最低耗时）数据波动大，当号段使用完之后还是会 hang 在更新数据库的 I/O 上，TP999 数据会出现偶尔的尖刺的问题，提供了双 buffer 优化。

### 双 buffer 优化

双 buffer 优化是为了 DB 取号段的过程能够做到无阻塞，不需要在DB取号段的时候阻塞请求线程，即当号段消费到某个点时就异步的把下一个号段加载到内存中，而不需要等到号段用尽的时候才去更新号段。

leaf 服务内部会有两个号段缓存区 buffer。当前号段已下发 10% 时，如果下一个号段未更新，则另启一个更新线程去更新下一个号段。当前号段全部下发完后，如果下个号段准备好了则切换到下个号段为当前 segment 接着下发，循环往复。

- 每个 biz_tag 都有消费速度监控
  - 通常推荐 segment 长度设置为服务高峰期发号 QPS 的 600 倍（10min），这样即使 DB 宕机，leaf 仍能持续发号 10-20min 不受影响
- 每次请求来临时都会判断下个号段的状态，从而更新此号段，所以偶尔的网络抖动不会影响下个号段的更新

仍然依赖 DB 的稳定性，需要采用主从备份的方式提高 DB的可用性，还有 Leaf-segment 方案生成的 ID 是趋势递增的，这样 ID 号是可被计算的，例如订单 ID 生成场景，通过订单 ID 号相减就能大致计算出公司一天的订单量，这个是无法忍受的。

## leaf-snowflake

完全沿用 [snowflake](#snowflake) 方案的 bit 位设计。

workerID(这里都是指数据中心 ID + 实例 ID 共 10bits) 的分配，当服务集群数量较小的情况下，完全可以手动配置；业务规模较大时引入了 Zookeeper 持久顺序节点的特性自动对 snowflake 节点配置 workerID。

启动步骤：

1. 启动 leaf-snowflake 服务，连接 Zookeeper，在 leaf_forever 父节点下检查自己是否已经注册过（是否有该顺序子节点）；
2. 如果有注册过直接取回自己的 workerID（zk 顺序节点生成的 int 类型 ID 号），启动服务；
3. 如果没有注册过，就在该父节点下面创建一个持久顺序节点，创建成功后取回顺序号当做自己的 workerID 号，启动服务。

为了减少对 Zookeeper 的依赖性，会在本机文件系统上缓存一个 workerID 文件。当 ZooKeeper 出现问题，恰好机器出现问题需要重启时，能保证服务能够正常启动。

针对 snowflake 的时钟回拨问题，通过校验自身系统时间与 leaf_forever/${self} 节点记录时间做比较然后启动报警的措施。

# Reference

[分布式系统-全局唯一ID实现方案](https://pdai.tech/md/arch/arch-z-id.html)

[Leaf——美团点评分布式ID生成系统](https://tech.meituan.com/2017/04/21/mt-leaf.html)

[分布式id生成器](https://books.studygolang.com/advanced-go-programming-book/ch6-cloud/ch6-01-dist-id.html)

[go 实现的 uuid](https://github.com/google/uuid)

[go 实现的 ulid](https://github.com/oklog/ulid)

[go 实现的 snowflake](https://github.com/bwmarrin/snowflake)
