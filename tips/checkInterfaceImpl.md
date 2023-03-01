
# 编译器检测类型是否实现接口

```go
type myWriter struct{}

// func (w myWriter) Write(p []byte) (n int, err error) {
// 	return
// }

func TestJudge(t *testing.T) {
	// 检查 *myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = (*myWriter)(nil)

	// 检查 myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = myWriter{}
}
```

上面的代码运行会报错，去掉 `Write()` 方法的注释后可以正常运行。

赋值语句会发生隐式地类型转换，在转换的过程中，编译器会检测等号右边的类型是否实现了等号左边接口所规定的函数。
