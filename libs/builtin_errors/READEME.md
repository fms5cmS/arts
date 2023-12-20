Go 内置库 errors

# Wrapper

Go 1.13 向 `errors` 和 `fmt` 包中引入了一种新特性用于简化 error(包含了另一种 error) 的使用。

[代码示例](./wrapper_test.go)

- 方法一：实现 `Wrapper` 接口

当 error e1 实现了 `Wrapper` 接口（即实现 `Unwrap()` 方法），则通过调用 `errors.Unwrap()` 方法即可得到 e1 所包装的错误。

一般标准的 error 都没有实现 Wrapper 接口，比如 `io.EOF`, 但是也有一小部分的 error 实现了，如 `os.PathError`，`os.LinkError`、`os.SyscallError`、`net.OpError`、`net.DNSConfigError` 等。

- 方法二：`fmt.Errorf()` + `%w`

更加常用的方式则是使用 `fmt.Errorf()` 和 `%w` 来进行包装，且该方法可以一次包装多个 error。

如果用到了 `%w`，则 `fmt.Errorf()` 返回的 error 自带 `Unwrap()` 方法，且 `Unwrap()` 方法的返回值即为 `%w` 对应的错误。

> 当一次性包装多个 error 时，则返回的 error 不会自带 `Unwrap()` 方法.

通过这两种方式，可以从最底层开始，逐层向上生成一个 error 树。

## errors.Unwrap()

通常在多层调用时，把最底层的 error 逐层包装传递上去，然后在最高层处理 error 的时候再调用 `errors.Unwrap()` 解开 error，逐层处理。

> `errors.Unwrap(xxErr)` 会返回 `xxErr.Unwrap()` 的结果，如果 xxErr 没有实现 Wrapper 接口，则返回 nil 

# errors.Is()

`errors.Is(err, target error)` 通过深度优先遍历的方式来检查整个 error 树中是否包含指定的目标 error，含义上更趋近于 Has。

[代码示例 TestIs](./isAs_test.go)

# errors.As()

`errors.As(err error, target any)` 通过深度优先遍历的方式来检查 error 树的每个 error，看能否将其中某个 error 赋值给目标变量。

[代码示例 TestAs](./isAs_test.go)

**常犯的一个错误就是：使用一个 error 类型的变量作为第二个参数**！见[代码示例 TestAsInvalid](./isAs_test.go)

# errors.join

项目中有时会处理多个 error：

```go
func (s *Server) Serve() error {
    var errs []error
    if err := s.init(); err != nil {
        errs = append(errs, err)
    }
    if err := s.start(); err != nil {
        errs = append(errs, err)
    }
    if err := s.stop(); err != nil {
        errs = append(errs, err)
    }
    if len(errs) > 0 {
        return fmt.Errorf("server error: %v", errs)
    }
    return nil
}
```

而在 Go1.20 中增加了 `errors.Join()` 来处理这种情况，它可以把多个 error 合并成一个 error！

然后就可以通过 `errors.Is()` 判断是否包含某个 error，或使用 `errors.As()` 来处理想要特殊处理的 error。

# error 的处理

见 [处理 error](../../review/errors/READEME.md)
