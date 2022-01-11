原文链接：[CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)

一般而言，如果类型 `T` 的方法与其指针类型 `*T` 相关联，要避免对 `T` 类型的对象的复制。

math/rand 实现了一个伪随机数生成器，如果不指定 seed 的话，生成的数是完全可预测的，即使使用 time.Now().UnixNano() 作为 seed，there are just a few bit of entropy(熵)；crypto/rand 则实现了一个加密的安全随机数生成器。如果要利用随机数生成器来生成 keys，建议使用 crypto/rand 包下的 Reader。见 [rand_test](./examples/rand_test.go)

空切片的声明：如果要生成空切片，推荐使用 `var s1 []string` 的方式，而不是 `s2 := []string{}`！

s1 和 s2 的 `len()`、`cap()` 返回的值都是 0，而 s1 是一个 `nil` 切片，s2 则不是！

非 nil 但长度为 0 的切片仅在有限的情况下优先使用，如：将对象序列化为 JSON 数据时，s1（`nil` 切片）会被转为 `null`，而 s2 则会被编码为 JSON 数组 `[]`。

Error 字符串不应有大写内容（除非开头是专有名词或首字母缩写词）或以标点结尾，因为它们通常是根据其他上下文打印的。如：使用 `fmt.Errorf("something bad")` 而不是 `fmt.Errorf("Something bad")`，这样在执行 `log.Printf("Reading %s: %v", filename, err)` 时中间的内容就不会有大写了。

生成 goroutine 时，请确认其何时或是否退出！阻塞 channel 的发送和接收数据的过程中可能会发生 goroutine 的泄露：即使被阻塞的 channel 无法访问了，垃圾回收器也不会终止 goroutine。

导入的包是按组来组织的，以空行进行分隔。标准库的通常位于第一组；

- `import _ "pkg"` 在导入时就调用了该包的 init 函数，且后续无法使用包的其他函数。这样的导入应为仅出现在程序的 main 包或需要它们的 tests 中！
- `import 别名 "pkg"` 给导入的包起别名，避免导入的包名冲突
- `import . "fmt"` 在调用 fmt 包中的函数时，不必再写包名即 `fmt.` 了，如可以直接使用 `Println("string")`
    - 一般不要使用这种导入方式，由于不确定代码中使用的是当前包还是被导入包中的顶级标识符（top-level identifier），代码可读性会很差
    - 这种导入方式一般用于 tests 中，下例中测试文件并未声明在 foo 包中，因为 `bar/testutil` 中同样导入了 foo ，所以使用 `import . foo` 来伪装该测文件是 foo 包的一部份（虽然实际该文件并未声明在 foo 包中）
    - 在 strings 包中就有这种使用！

标识符命名时，专有名词要保持一致大小写而不要首字母缩写。如写 URL 或 url （URLPony 或 urlPony） 而不是 Url、写 ServeHTTP 而不写 ServeHttp、写 xmlHTTPRequest 或 XMLHTTPRequest、写 appID 代替 appId。

不要仅仅为了避免在函数内声明变量就对结果参数命名。尽量不要对结果参数命名！

- 方法接收者命名：一般使用其所属类型的一两个字母缩写即可。
- 方法接收者类型：如有疑问优先使用指针接收者，但是有时值接收者更好，如小型不变的结构体或基础类型的值。
    - 如果接收者是 map、func 或 chan，不要使用指针接收者；
    - 如果接收者是 slice，且方法并未 reslice 或重新分配 slice，不要使用指针接收者；
    - 如果方法需要修改（mutate）接收者，那就必须使用指针接收者；
    - 如果接收是包含了 `sync.Mutex` 或类似同步原语的结构体，必须使用指针接收者以避免复制；
    - 如果接收者是大结构体或数组，指针接收者更有效；
    - 根据函数或方法是否会使接收者变化来判断；
        - 方法被调用时，值类型接收者会复制接收者，因此外部更新不会作用域接收者
        - 如果变更要在原指针接收者中可见，那就必须是指针接收者
    - 如果接收者是结构体、数组或 slice，且它的任何元素都是指向可能变更的内容，那就更倾向于指针接收者，可读性会更好
    - 如果接收者是小的数组或结构体（两者都是值类型），没有可变字段或指针，或只是一个简单的基础类型（如 `int` 或 `string`），值接收者更有效
    - 值接收者可以减少生成的垃圾，如果将值传递给值方法，可以使用堆栈上的副本而不是重新在堆上分配（编译器会尽量避免这种分配，但并非总是成功）。对于这种情况，先进行分析（profiling）再决定选择哪种接收者。

tests 失败应该返回有帮助的信息（错误、输入、实际得结果、期望结果分别是什么），如下例，注意，这里的顺序是 `实际 != 期望`，且信息中使用的也是这样的顺序。也可以写成表驱动测试。

```go
if got != tt.want {
  // // or Fatalf, if test can't test anything more past this point
  t.Errorf("Foo(%q) = %d; want %d", tt.in, got, tt.want) 
}
```
