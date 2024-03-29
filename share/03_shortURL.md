
短链接字符少，便于发布和传播（微博等应用会限制发布字数，手机发送短信时每 70 个汉字按一条短信收费等）。

短链接访问流程：
1. 浏览器输入短链接；
2. DNS 解析到域名对应短网址服务的 IP 后向其发送 GET 请求查询短链接域名后面的短码
3. 短网址服务器通过短码获得对应的长 URL
4. 请求重定向转到对应的长 URL

# 短网址生成

通常有两种算法：
- 自增序列算法：
  - 容易理解，永不重复；
  - 短码长度不固定，随 id 增长会递增，如果要长度固定，可以让 id 从指定数字开始自增；过度依赖数据库导致性能比较差
- 摘要算法：
  - 短码长度固定
  - 存在碰撞（重复）的概率

## 自增序列算法

**自增 ID + 62 进制编码**，主要利用低进制转高进制时，字符数会减少的特征。

假设短网址服务器的域名为 xx.cn，当长网址转短网址请求过来后，先利用发号器生成一个自增的 ID，如 `10000`，然后将这个自增 ID 转为 62 进制后得到 `2Bi`，这个长网址对应的短网址就是 xx.cn/2Bi。

### 为什么使用 62 进制转换？

因为 62 进制转换后只含数字+小写+大写字母。而 64 进制转换会含有 `/`，`+` 这样的符号（不符合正常URL的字符）。

仅 6 位的 62 进制就已经有 62^6 = 560 亿个组合。

## 摘要算法

见 examples/shortURL/summary_test.go

# Questions

Q1：短网址和长网址是一对一还是一对多？

一对多！如果一个长网址与一个短网址一一对应，那么在数据库中，仅有一行数据，无法区分不同的来源，就无法做数据分析了。

一般而言，一个长网址，在不同的地点，不同的用户等情况下，生成的短网址应该不一样，这样，在后端数据库中，可以更好的进行数据分析。如，以一个六位的短网址作为唯一 ID，这个 ID 下可以挂各种信息，如用户名、所在网站等，收集这些信息后，再做大数据分析，挖掘数据价值。

Q2：301 重定向 or 302 重定向？

301 是永久重定向，302 是临时重定向。短地址一经生成就不会变化，所以用 301 是符合 http 语义的。

但是，如果用了 301，搜索引擎搜索时会直接展示真实地址，那么短网址服务就无法统计到短网址被点击到的频率了，也无法收集用户的Cookie, User Agent 等信息，所以短网址服务商通常会使用 302 重定向。

Q3：短时间向服务器发送大量请求，迅速耗光 ID 怎么办？

限制 IP 的单日请求数；使用 Redis 混存，存储 长网址 -> ID，仅存储一天内的数据，使用 LRU 机制进行淘汰。

即使大量发同一个长网址过来，也可以从缓存服务器直接返回短网址，而不会每次都达到数据库层。
