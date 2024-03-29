
共识重点在于达成一致的过程或方法，一致性问题在于最终对外表现的结果。

# Consistency

一致性（Consistency）在分布式系统领域中是指对于多个服务节点，给定一系列操作，在约定协议的保障下，使得多个节点对外呈现的数据的状态一致。

> 注：一致性关注的是系统呈现的状态，并不关注结果是否正确；例如，所有节点都对某请求达成否定状态也是一种一致性。

实际系统的一致性主要包括以下两大类：

- 强一致性（Strong Consistency）
    - 线性一致性（Linearizability Consistency）
    - 顺序一致性（Sequential Consistency）

> 线性一致性，要求整个系统所有节点上所有线程的同一类操作必须保证进行全局唯一排序，所以通常需要依赖于一个全局的时钟或锁
>
> 顺序一致性，要求每个进程自身操作的顺序（local order）跟实际请求顺序一致，而在不同进程之间的执行顺序则不要求

- 弱一致性（Weak Consistency）
    - 最终一致性（Eventual Consistency）

# Consensus

共识（Consensus）是指在分布式系统中多个节点对某件事（如多个事务请求先执行谁）达成一致看法的过程。因此，达成某种共识并不意味着保障了一致性。

> 实践中，要保障系统满足不同程度的一致性，往往需要通过共识算法来达成。
>
> 对于分布式系统而言，各个节点通常都是相同的确定性状态机模型（也叫状态机复制问题，State-Machine
> Replication），从相同初始状态开始接收相同顺序的指令，则可以保证相同的结果状态，因此，系统中多个节点最关键的是对多个事件的顺序达成共识，，即排序。

达成共识要解决两个基本问题：

1. 如何提出一个待共识的提案？如通过令牌传递、随机选取、权重比较、求解难题等
2. 如何让多个节点对该提案达成共识（统一或拒绝）？如投票、规则验证等

而在实际中，不同节点间会出现各种问题，如通信网络中断、节点故障、被入侵节点伪造信息等，破坏正常的共识过程。一般，把出现故障但不会伪造信息的情况称为“非拜占庭错误（Non-Byzantine Fault）”或“故障错误（Crash Fault）”；而伪造信息恶意响应的情况称为“拜占庭错误（Byzantine Fault）”，其对应的节点为拜占庭节点。

根据解决的场景是否允许拜占庭错误情况，共识算法可以分为 **CFT(Crash Fault Tolerance，崩溃容错)** 和 **BFT(Byzantine Fault Tolerance，拜占庭容错)** 两类。

对于非拜占庭错误的情况，有很多经典的算法，如 Paxos、Raft 及其变种等，这类容错算法往往性能比较好，处理较快，容忍不超过一般的故障节点。

对于要容忍拜占庭错误的情况，包括 PBFT(Practical Byzantine Fault Tolerance) 为代表的确定性算法、PoW 为代表的概率性算法。

> 确定性算法一旦达成共识就不可逆转，即共识是最终结果；而概率类算法的共识结果则是临时的，随着时间推移或某种强化，共识结果被推翻的概率越来越小，最终成为事实上结果。拜占庭类容错算法往往性能较差，容忍不超过 1/3 的故障节点。

## FLP 不可能原理

FLP 不可能原理：在网络可靠、但允许节点失效的最小化**异步模型系统**中，**不存在**可以解决一致性问题的确定性共识算法！

[分布式系统的同步和异步](./READEME.md#同步异步)

尽管 FLP 不可能原理说明了不存在完美的共识方案，但是在实际中，付出一定代价的情况下，也是可以达成一定程度共识的。即 CAP 理论。

## CAP 理论

CAP 理论：分布式系统不可能同时满足 一致性（Consistency）、可用性（Availability）、分区容错性（Partition tolerance）三个需求，设计中往往需要弱化对某个特性的需求。

- 一致性（Consistency），在分布式系统中的所有数据备份，在同一时刻是否同样的值（等同于所有节点访问同一份最新的数据副本）；
- 可用性（Availability），在集群中一部分节点故障后，集群整体是否还能响应客户端的读写请求（对数据更新具备高可用性）；
- 分区容错性（Partition tolerance），以实际效果而言，分区相当于对通信的时限要求。系统如果不能在时限内达成数据一致性，就意味着发生了分区的情况。

> 网络本身不可能做到 100% 可靠，所以分区是一个必然的现象，也因此 P 要素是必须选择的，分布式系统理论上不可能选择 CA 系统。

应用场景：

- AP 系统，弱化一致性：对结果一致性不敏感的应用，可以允许在新版本上线后过一段时间才最终更新成功，期间不保证一致性
    - 如网站静态页面内容、实时性较弱的查询类数据库等，简单分布式同步协议如 Gossip，以及 CouchDB、Cassandra 数据库等，都为此设计
- CP 系统，弱化可用性：对结果一致性很敏感的应用
    - 如银行取款机，当系统故障时候会拒绝服务。MongoDB、Redis、MapReduce 等为此设计
    - Paxos、Raft 等共识算法，主要处理这种情况。在 Paxos 类算法中，可能存在着无法提供可用结果的情形，同时允许少数节点离线
- AC 系统，弱化分区容错性
    - 现实中网络分区出现概率较小，但很难完全避免
    - 两阶段的提交算法，某些关系型数据库以及 ZooKeeper 主要考虑了这种设计
    - 实践中，网络可以通过双通道等机制增强可靠性，实现高稳定的网络通信。

## 基于事务过程的系统的一致性

[ACID 原则](./04_transaction.md)描述了分布式数据库需要满足的一致性需求，同时允许付出可用性的代价。

[BASE 原则](./04_transaction.md#分布式事务)面向大型高可用分布式系统，主张牺牲强一致性，而实现最终一致性，来换取一定的可用性。

## Paxos&Raft&ZAB算法

[分布式事务](./04_transaction.md#分布式事务)的 [2PC](./04_transaction.md#2pc)、[3PC](./04_transaction.md#3pc) 都只是一定程度上缓解了提交冲突的问题，而无法确保系统的一致性。首个有效的共识算法是后来提出的 Paxos 算法。

Paxos 问题指分布式的系统中存在故障（crash fault），但不存在恶意（corrupt）节点的场景（即可能消息丢失或重复，但无错误消息）下如何达成共识，即非拜占庭错误的情况。

### Paxos 算法

原理类似于 [2PC](./04_transaction.md#2pc) 算法，通过消息传递来逐步消除系统中的不确定状态。该算法中存在三种逻辑角色的节点，在实现中同一节点可以担任多个角色：

- Proposer(提案者)：提出提案，等待大家批准（Chosen）或决议（Value）
    - 系统中提案都拥有一个自增的唯一提案号。往往由客户端担任该角色
- Acceptor(接受者)：负责对提案进行投票，接收（Accept）提案
    - 往往由服务端担任该角色
- Learner(学习者)：获取批准结果，并帮忙传播，不参与投票过程
    - 可为客户端或服务端

多个提案者先要争取到提案的权利（得到大多数接受者的支持）；成功的提案者发送提案给所有人进行确认，得到大部分人确认的提案称为批准的决议。

Paxos 能保证在超过一半的节点正常工作时，系统总能以较大概率达成共识。

> 算法需要满足安全性（Safety）和存活性（Liveness）两方面的约束要求：
>
> Safety：保证决议（Value）结果是对的，无歧义的，不会出现错误情况
>
> 在一次执行中只批准（Chosen）一个最终决议。被多数接受的结果成为决议。
>
> Liveness：保证决议过程能在有限时间内完成

- 单提案者+多接受者

提案者只要收到来自多数接受者的投票，即可认为通过。容易发生单点故障，一旦提案者故障，则整个系统无法工作。

- 多提案者+单接受者

接受者收到多个提案，选第一个提案作为决议，发送给其他提案者即可。容易发生单点故障。

- 多提案者+多接受者

情况一：同一时间段（如一个提案周期）内只有一个提案者，此时会退化到单提案者的情形。需要设计一种机制保障提案者的正确产生，如按照时间、序列等方式。

情况二：允许同一时间段内出现多个提案者，那么同一个节点可能会收到多份提案，如何对其区分呢？提案时带上不同的序号。节点根据序号判断接受哪个提案，通常采用递增序号，选择接受序号最大的提案，因为旧提案可能基于过期数据，导致失败概率更大。

同时允许多个提案，意味着很可能单个提案人无法集齐足够多的投票；

另一方面，提案者发出提案申请之后，会收到来自接受者的反馈。一种结果是提案被大多数接受者接受了，一种结果是没被接受。没被接受的话，可以过会再重试。即便收到来自大多数接受者的答复，也不能认为就最终确认了。因为这些接受者自己并不知道自己刚答复的提案可以构成大多数的一致意见。所以需要引入一个新的阶段，即提案者在第一阶段拿到所有反馈后，需要再次判断这个提案是否得到大多数的支持，如果支持则需要对其进行最终确认。

#### 两阶段提交

注：Paxos 并不一定能保证每一轮都能提交提案。

- 准备阶段（Prepare）：通过锁来解决对哪个提案内容进行确认的问题

提案者向多个接受者发送计划提交的提案编号 n，试探是否可以锁定多数接受者的支持；

接受者 i 收到提案编号 n，检查回复过的提案的最大编号 Mi，如果 n > Mi，则向提案者返回准备接受提交的最大编号的提案 Pi，并不再接受小于 n 的提案，同时更新 Mi = n，用于让接受者筛选出它收到的最大编号的提案，接下来只接受其后续提交。

- 提交阶段（Commit）：决绝大多数确认最终值的问题

某个提案者如果收到大多数接受者的回复，则准备发出带有 n 的提交消息，如果收到的回复中带有提案 Pi（说明自己看到的信息过期），则替换选编号最大的 Pi 的值作为提案值，否则指定一个新的提案值，如果没收到多数回复，则再次发出请求；

接受者 i 收到序号 n 的提交消息，如果发现 n >= Pi，则接受提案，并更新 Pi 序号为 n。

### Raft 算法

Raft 算法通过先选出领导节点来简化流程和提高效率。实现上解耦了领导者选举、日志复制和安全方面的需求，并通过约束减少了不确定性的状态空间。

算法包括三种角色：

- Leader(领导者)：主节点，同一时刻只有一个 Leader，负责协调和管理其他节点
- Candidate(候选者)：每个节点都可以称为 Candidate，节点在该角色下才可以被选为新的 Leader
- Follower(跟随者)：不可以发起选举

每个任期内选举一个全局的 Leader，Leader 负责从客户端接收请求，并分发到其他节点，并决定日志 log 的提交。

其过程包含两个主要阶段：

1. 领导者选举（term）：
    1. 一开始所有节点都是 Follower，在随机超时发生后还未收到来自 Leader 或 Candidate 消息则从 Follower 转为 Candidate，并向其他节点发送选举请求
    2. 其他节点根据收到的选举请求的先后顺序，回复是否同意为 Leader，注：每一轮选举中，一个节点只能投出一票
    3. 得票超过一半者成为 Leader，其他节点状态由 Candidate 转为 Follower
    4. Leader 和 Follower 间会定期发送心跳，以检测 Leader 是否存活
2. 同步日志：
    1. Leader 决定系统中最新的日志记录，并强制所有的 Follower 来刷新到这个记录
    2. 数据的同步是单向的（Leader 向 Follower），确保所有节点看到的视图一致

Raft 算法中，领导者选举是周期进行的；而如果 Follower 如果发现心跳超时未收到，则可以认为 Leader 下线，会立即尝试发起新的选举过程。

### ZAB 算法

ZAB（ZooKeeper Atomic Broadcast）选举算法是为 ZooKeeper 实现分布式协调功能而设计的。相较于 Raft 算法的投票机制，ZAB 算法增加了通过节点 ID 和数据 ID 作为参考进行选主，节点 ID 和数据 ID 越大，表示数据越新，优先成为主。

算法包括三中角色：

- Leader
- Follower
- Observer，无投票权

选举过程中，每个节点有四个状态：

- Looking：选举状态，节点处于该状态时，它会认为当前集群中没有 Leader，因此自己进入选举状态
- Leading：领导者状态，集群中已经选择出 Leader，且当前节点为 Leader
- Following：跟随者状态，集群中已经选择出 Leader，所有非 Leader 节点状态为 Following
- Observing：观察者状态，当前节点为 Observer，没有投票权和选举权

投票过程中，每个节点都以一个唯一的三元组 `(server_id,server_zxID,epoch)`，分别表示本节点的唯一 id，本节点存放的数据 id（越大代表数据越新，选举权重越大），当前选举轮数（一半用逻辑时钟表示）。

选举过程中通过 `(vote_id,vote_zxID)` 来表明投票给哪个节点，其中 vote_id 表示被投票节点的 ID，vote_zxID 表示被投票节点的服务器 zxID。

ZAB 算法选主的原则是：server_zxID 最大者成为 Leader；若 server_zxID 相同，则 server_id 最大者成为 Leader。

## BFT 类共识（确定性算法）

> BFT 类共识无法支持大规模节点，复杂度会很高，更适用于联盟链！

拜占庭问题（Byzantine Problem）也叫拜占庭将军（Byzantine Generals Problem）问题，讨论的是在少数节点有可能作恶（消息可能被伪造）的场景下，如何达成共识。

拜占庭容错（Byzantine Fault Tolerant，BFT）讨论的是容忍拜占庭错误的共识算法。解决的是在网络通信可靠，但节点可能故障和作恶的情况下如何达成共识。

拜占庭问题的解决方案一致都存在运行过慢或复杂度过高的问题，直到实用拜占庭容错（Practical Byzantine Fault Tolerance，PBFT）算法的提出。

> PBFT 首次将 BFT 算法复杂度从指数级降低到了多项式（平方）级，其可以在恶意节点不超过总数 1/3 的情况下同时保证 Safety 和
> Liveness。

PBFT 算法采用密码学相关技术（RSA 签名算法、消息验证编码和摘要）确保消息传递过程无法被篡改和破坏。其基本过程：

1. 通过轮换或随机算法选出某个节点为主节点，此后只要主节点不切换，则称为一个视图 View
2. 在某个 View 中，客户端将请求 `<REQUEST,operation,timestammp,client>` 发送到主节点，主节点负责广播请求到的所有其他从节点并完成共识
3. 所有节点处理完请求，将处理结果 `<REPLY,view,timestamp,client,id_node,response>` 返回给客户端
4. 客户端检查是否收到至少 f+1 个来自不同节点的相同结果，将其作为最终结果

f 代表了 Byzantine 节点数量。

主节点广播过程：

1. Pre-Prepare(预准备)：对于从客户端收到的请求，主节点对其分配提案编号，然后发出预准备消息 `<PRE-PREPARE,view,n,digest>` 给各个从节点，主节点需要对预准备消息进行签名。这是为请求分配序号 n 并通知其他节点，因此可以不包含原始的请求消息，可通过其他方式将请求同步到从节点。
2. Prepare(准备)：从节点收到预准备消息后，检查消息（核对签名、视图、编号），如果消息合法，则向其他节点发送准备消息 `<PREPARE,view,n,digest,id>` 带上自己的 id 信息，并添加签名，收到准备消息的节点同样需要对消息进行合法性检查。节点集齐至少 2f+1 个验证过的消息则认为验证通过，把这个准备消息写入本体提交消息日志中。这是为了确认大多数节点已经对序号达成共识。
3. Commit(提交)：主节点广播 commit 消息 `<COMMIT,v,n,digest,id>` 并添加自己的签名，告诉其他节点某个编号为 n 的提案在视图 v 中已经处于提交状态。如果集齐至少 2f+1 个验证过的 commit 消息，则说明提案被整个系统接受。

PBFT 算法和 Raft 算法过程类似，但是 PBFT 中没有假设主节点一定可靠，因此增加了额外的从节点间的交互，当发现主节点不可靠时通过重新选举选出新的主节点。

具体实现上还包括 checkpoint（同步节点状态和清理本地日志数据）、视图切换（重新选举主节点）等机制。

拜占庭问题之所以难解，在于任何时候系统中都可能存在多个提案（因为提案成本很低），并且在大规模场景下要完成最终确认的过程容易受干扰，难以达成共识。

[Practical Byzantine Fault Tolerance](https://pmg.csail.mit.edu/papers/osdi99.pdf)

## PoX 类共识（概率性算法）

> PoX 类共识更适用于公链

比特币网络的 PoW（Proof of Work，工作量证明）的概率型算法思路，从两个角度解决大规模场景下的拜占庭容错问题：

1. 限制一段时间内整个网络中出现提案的个数（通过工作量证明来增加提案成本）
2. 丢掉最终确认的约束，约定好始终沿着已知最长的链进行拓展

共识的最终确认是概率意义上的存在。这样，即便有人试图恶意破坏，也会付出相应的经济代价（超过整体系统一半的工作量）。后来的各种 PoX 系列算法，也都是沿着这个思路进行改进，采用经济博弈来制约攻击者。

PoW 机制的缺点也很明显，共识达成的周期长、效率低，资源消耗大。

- PoS(Proof of Stake, 权益证明)，由系统权益代替算力来决定区块记账权，拥有的权益越大获得记账权的概率就越大。

PoS 算法类似于现实生活中的股东机制，拥有股份越多的人越容易获取记账权。持币越多的人就越容易挖到区块并得到激励，而持币少的人基本没有机会，这样整个系统的安全性实际上会被持币数量较大的一部分人掌握，容易出现垄断现象。

- DPoS(Delegated Proof of Stake, 委托权益证明)

类似股份制公司的董事会制度，普通股民虽然拥有股权，但进不了董事会，他们可以投票选举代表（受托人）代他们做决策。DPoS 是由被社区选举的可信帐户（受托人，比如得票数排行前 101 位）来拥有记账权。

- PoC/PoS(Proof of Capacity / Storage, 容量/存储证明)

### 无利害攻击

在 PoS 共识机制中存在一个较为普遍的问题——无利害攻击。

PoS 验证者可以选择同时在两条链上进行投票，不论最后哪条链成为“最长链”胜出，验证节点都可以获得区块奖励，而随着时间推移，这种无成本为多条链出块的投票行为会助长区块链分叉行为。

### Casper 共识

ETH 2.0 会采用基于 PoS 的 Casper 共识机制：

- 质押机制：成为 ETH2.0 验证者需要至少质押 32ETH 作为保证金
- 随机选举：Casper 验证节点委员会**每隔一个周期进行重新选举轮换**，随机指派验证节点负责指定分片内的区块校验，避免分片内验证节点形成“合谋”
-
惩罚机制：验证者需要评估其他节点投注的区块，确保投注给大概率胜出的区块，如果验证节点投注其他区块将无法获得奖励，如果投注后又改投其他区块将被没收保证金，消极投注如离线行为，也将可能被处罚保证金。可以有效规避[无利害攻击](#无利害攻击)问题
