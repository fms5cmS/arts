
当要构建的对象很大，且需要多个步骤时，使用构造器模式，减少构造函数的大小。

[代码示例](../examples/designPatterns/builder_test.go)，还有其中的 error 处理方式：将原本在外部调用时的错误判断，分散到了每个步骤中。

Go 中还有一种最佳实践是 [Option 模式](../examples/designPatterns/option_test.go)，可以类比


