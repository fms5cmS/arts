| 目录                | 描述                   | 目的               |
|-------------------|----------------------|------------------|
| algorithm         | 每周至少一道 Leetcode 的算法题 | 编程训练和学习          |
| [review](#review) | 阅读并点评至少一篇英文技术文章      | 学习英文             |
| [tip](#tips)      | 学习至少一个技术技巧           | 总结和归纳日常生活中遇到的知识点 |
| [share](#share)   | 分享一篇有观点和思考的技术文章      | 建立自己的影响力，能够输出价值观 |

补充：

| 目录                  | 描述           |
|---------------------|--------------|
| [biz](#biz)         | 业务相关知识记录，无代码 |
| scripts             | 脚本记录（todo）   |
| [summary](#summary) | 代码汇总         |
| [books](#books)     | 书籍阅读         |
| [libs](#libs)       | 包阅读          |

[零散的记录](some.md)

# share

- [分布式锁](share/01_distributedLock.md)
- [限流](share/02_rateLimit.md)
- [事务](share/04_transaction.md)：MySQL 事务、Redis 事务、分布式事务
- [持久化](share/05_persistence.md)：MySQL 事务的持久性、Redis 持久化、
- [哈希](share/06_hash.md)：内存结构&序列化方案、Redis 中的哈希表、Golang 的 map
- [索引](share/07_index.md)：索引的常用数据类型、MySQl 的索引
- [锁](share/08_lock.md)：死锁、分布式锁、MySQL 的锁、Golang 的锁
- [一致性&共识](share/12_consistency&consensus.md)
- [池化技术](share/15_pool.md)

业务相关内容：

- [短链接](share/03_shortURL.md)
- [如何存储一棵树](share/11_storeTree.md)
- [唯一标识](share/14_uniqueIdentifier.md)

较宽泛内容（了解）：

- [加密](share/09_cryptology.md)
- [微服务笔记](share/10_SomeMicroServiceNote.md)
- [服务发现](share/13_serviceDiscovery.md)

# biz

- [撮合引擎](biz/01_matchEngine.md)
- [区块链-借贷业务](biz/02_blockchain-loan.md)
- [区块链-dex](biz/03_blockchain-dex.md)

# libs

- [asynq 分布式任务队列](libs/asynq/README.md) 
- [内置 errors 包](libs/builtin_errors/READEME.md) 
- [内置 sync 包](libs/builtin_sync/REDEME.md) 含 sync/atomic 包
- [内置 errors 库](libs/builtin_errors/READEME.md) 

# review

- [go 代码 review 建议](review/codeReviewComments/README.md)
- [go 并发模型](review/concurrencyPatterns/README.md)
- [第三方 errors 库](review/errors/READEME.md)
- [如果对齐的内存写入是原子的，为什么还需要 sync/atomic 包](review/whyNeedAtomicPackage/README.md)
- [Go 读取大文件](review/readFile/README.md)

# summary

- [基于以太坊基础 api 的 go 访问接口](summary/blockchain/README.md)
- [MD5 和 AES 加密使用示例](summary/encrypt/aes.go)
- [接入聊天工具](summary/msg/README.md)
- [读取文件的部分操作](summary/readFile/README.md)

# tips

- [go 编程中的一些建议](tips/README.md)
- 设计模式（go 示例）
- [go 检查接口实现](tips/checkInterfaceImpl.md)
- [rand 包示例](tips/rand.md)：根据字符源生成指定长度的字符串
- [time 包](tips/time.md)
- [unsafe 包](tips/unsafe.md)

# books

- [支付方法论](books/payment/支付方法论.md)