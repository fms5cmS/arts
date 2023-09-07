
# 根据字符源生成指定长度的字符串

如生成验证码的业务。

注：这里使用 rand 包是 golang.org/x/exp/rand 而非 math/rand，但思想是一致的，仅细节不同。

通常会使用 `rand.Intn(length)` 生成随机索引，然后取字符源中该位字符，多次循环调用从而获得指定长度的结果。示例见 [randStr](examples/randTip/randStr.go).RandStrSimple 函数。

这种方案会多次调用 `rand.Intn()` 函数，该函数底层是调用的 `rand.Uint64()`，而一次该操作可以得到一个 64 位的数字，那么是否可以调用一次 `rand.Uint64()` 然后对其进行位运算从而得到多个随机值呢？

这是可行的，思路如下：

假设字符源 chars = "0123456789"，其长度为 10，所以在生成随机索引时，索引最大值为 10，对应二进制为 1010，也就是说一个四位的二进制数就可以代表一个索引了，而一次 `rand.Uint64()` 得到的就是一个 64 位的二进制数，里面包含了 64/4=16 个随机索引，这就大大减少了 `rand.Uint64()` 的调用次数，性能也会提升很多。

代码实现见 [randStr](examples/randTip/randStr.go).RandStrQuick 函数。

> 注：一个四位的二进制数最大是 1111 对应十进制的 15，这大于 chars 的长度，所以在处理时，需要判断这种特殊情况。

在代码目录下使用 `go test -benchmem -bench=.` 来对比两种方案的[性能测试](examples/randTip/randStr_test.go)。
