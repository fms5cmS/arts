
Reference:

- 可以整行写入内存

[The One Billion Row Challenge in Go: from 1m45s to 3.4s in nine solutions](https://benhoyt.com/writings/go-1brc/)

[The One Billion Row Challenge](https://www.morling.dev/blog/one-billion-row-challenge/)

[Code](https://github.com/benhoyt/go-1brc/tree/master)

[4 秒处理 10 亿行数据！ Go 语言的 9 大代码方案，一个比一个快](https://mp.weixin.qq.com/s/iylZAKZfxLL6SYruww_8zA)

- 整行不一定能写入内存

[Reading 16GB File in Seconds, Golang](https://medium.com/swlh/processing-16gb-file-in-seconds-go-lang-3982c235dfa2)

solutions:

[read1](./read1.go)，最常见最简单的方案，使用 bufio.Scanner 来读取行数据

[read2](./read2.go)，针对 read1 中重复处理的数据，进行优化，减少处理次数

[read3](./read3.go)，根据文件中的实际数据格式优化之前的数据处理函数

[read4](./read4.go)，根据文件中的实际数据格式优化数据处理函数，**整数的处理比浮点数处理性能更好**

[read5](./read5.go)，根据文件中的实际数据格式优化之前的数据处理函数

[read6](./read6.go)，去掉 bufio.Scanner

[read7](./read7.go)，对数据的处理使用自定义哈希表，但是需要考虑哈希冲突、桶预分配、分配不够还需要如何调整哈希表大小的逻辑

[read8](./read8.go)，在 read1 的基础上，进行并行处理

[read9](./read9.go)，在 2～7 的所有基础上进行并行处理

注意：所有针对数据格式的优化都需要根据实际遇到的数据调整！

一行文件太大时，无法一次性写入内存，可以分片处理。

[readBigLine](./readBigLine.go)，读取不含换行符的大文件，使用 bufio.NewReader() 将文件分块加载到内存中！
