
[检测结构体是否实现了某个接口](./checkInterfaceImpl.md)

[time 包的注意事项](./time.md)

[unsafe 包的基础使用](./unsafe.md)

# Go 编程中的一些建议

- 如果需要把数字转换成字符串，使用 `strconv.Itoa()` 比 `fmt.Sprintf()` 要快一倍左右。
- 尽可能避免把 string 转成 `[]byte` ，这个转换会导致性能下降，可以[借助 unsafe 包来实现零拷贝转换](./examples/unsafe-test/example2_test.go);
- 如果在 for-loop 里对某个 Slice 使用 `append()`，请先把 Slice 的容量扩充到位，这样可以避免内存重新分配以及系统自动按 2 的 N 次方幂进行扩展但又用不到的情况，从而避免浪费内存；
- 使用 `StringBuffer` 或是 `StringBuild` 来拼接字符串，性能会比使用 `+` 或 `+=` 高三到四个数量级；
- 尽可能使用并发的 goroutine，然后使用 `sync.WaitGroup` 来同步分片操作；
- 避免在热代码中进行内存分配，这样会导致 gc 很忙。尽可能使用 `sync.Pool` 来重用对象；
- 使用 lock-free 的操作，避免使用 mutex，尽可能使用 `sync/Atomic` 包
  - 无锁队列实现：https://coolshell.cn/articles/8239.html
  - 无锁 HashMap 实现：https://coolshell.cn/articles/9703.html
- 使用 I/O 缓冲，I/O 是个非常非常慢的操作，使用 `bufio.NewWrite()` 和 `bufio.NewReader()` 可以带来更高的性能；
- 对于在 for-loop 里的固定的正则表达式，一定要使用 `regexp.Compile()` 编译正则表达式。性能会提升两个数量级；
- 如果你需要更高性能的协议，就要考虑使用 protobuf 或 msgp 而不是 JSON，因为 JSON 的序列化和反序列化里使用了反射；
- 在使用 Map 的时候，使用整型的 key 会比字符串的要快，因为整型比较比字符串比较要快。


