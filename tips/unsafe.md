
unsafe 包提供了两个能力：
1. 任何类型的指针和 unsafe.Pointer 可以相互转换
2. uintptr 类型和 unsafe.Pointer 可以相互转换

> uintptr 没有指针的语义，因此 uintptr 指向的对象会被 gc 回收；而 unsafe.Pointer 有指针语义，可以保护它指向的对象在"有用"的时候不会被 gc。

- [利用 unsafe 修改私有成员](./examples/unsafe-test/example1_test.go)

- [实现字符串和 byte 切片的零拷贝转换](./examples/unsafe-test/example2_test.go)
