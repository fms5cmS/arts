# Golang 编译

## go build

- 构建：`go build [参数]`
  - 命令源码文件及其依赖的库源码文件，并生成一个可执行的文件
  - `-a` 将命令/库源码文件全部重新构建
  - `-n` 把编译器涉及的命令全部打印出来，但不会真的执行，方便学习
  - `-race` 开启竞态条件检测，支持的平台有限制
  - `-x` 打印编译期间用到的命名，与 `-n` 的区别在于它不仅打印还会执行
  - `-ldflags “-X '包名.变量名=值'”` 向程序中指定包的变量传递值
  - `-o 文件名` 指定编译后输出的可执行文件名
  - `-gcflags '-N -l'` 禁止内联
    - `-N` 禁止编译优化
    - `-l` 禁止内联，可以在一定程度上减小可执行程序大小

使用 `go build` 生产可执行文件后，`./可执行文件名` 即可运行，使用

新建 first.go 文件：

```go
package main
import "fmt"
func main() {
  a := 1 + 2
  b := 10
  c := a * b
  fmt.Println(c)
}
```

执行 `go build -n first.go`，查看编译器的命令可以发现编译的核心是通过 `compile`、`buildid`、`link` 三个命令导出可执行文件 a.out，然后通过 `mv` 将其移动到当前文件夹下，并修改为和项目文件一样的名字。

## Go 编译器

原文链接：[Go 语言编译过程概述](https://mp.weixin.qq.com/s/lhle7ahjP7g9GIBvK8J23A)

Go 语言编译器的源代码在 cmd/compile 目录中，目录下的文件共同构成了 Go 语言的编译器。

编译原理中编译分为：

1. 编译前端：词法分析 -> 语法分析 -> 语义分析
2. 编译后端：中间码生成 -> 代码优化 -> 机器码生成

Go 的编译器在逻辑上可以被分成四个阶段：词法与语法分析、类型检查和 AST 转换、通用 SSA 生成和最后的机器代码生成

- 通过词法分析器(Lexer)对源码文件进行词法分析(Lexical Analysis)
  - 根据构词规则将源代码中的字符串序列转为为 Token 序列
  - Token 可以分多个类型：变量名、字面量、操作符、分隔符、关键字等
- 通过语法分析器(Parser) 进行语法分析(Syntax analysis 或 Parsing)
  - 将 Token 序列转为可识别的程序语法结构，会生成一棵 AST(抽象语法树)
  - 如果在语法解析的过程中发生了任何语法错误，都会被语法解析器发现并将消息打印到标准输出上，整个编译过程也会随着错误的出现而被中止。
  - 每一个 AST 都对应着一个单独的 Go 语言文件，这个 AST 中包括当前文件属于的包名、定义的常量、结构体和函数等。
- 类型检查
  - 对语法树中定义和使用的类型进行检查，类型检查分别会按照顺序对不同类型的节点进行验证，按照以下的顺序进行处理：
    - 常量、类型和函数名及类型；
    - 变量的赋值和初始化；
    - 函数和闭包的主体；
    - 哈希键值对的类型；
    - 导入函数体；
    - 外部的声明；
  - 所有的类型错误和不匹配都会在这一个阶段被发现和暴露出来。
  - 不止会对树状结构的节点进行验证，同时也会对一些内建的函数进行展开和改写，例如 `make` 关键字在这个阶段会根据子树的结构被替换成 `makeslice` 或者 `makechan` 等函数。

到了这里就说明代码结构、语法没有问题了。

- 中间代码生成
  - Go 编译器的中间代码使用了 SSA(Static Single Assignment Form，静态单赋值)的特性
    - 如果一个中间代码具有静态单赋值的特性，那么每个变量就只会被赋值一次
    - 其主要作用是代码的优化

```go
// 下面这段代码中，y1 的值和 x1 的值是完全没关系的，所以在机器码生成时其实可以省略第一步，这样就能减少需要执行的指令来优化这一段代码。
x1 := 1
x2 := 2
y1 := x2
```

- 机器码生成
  - 经过优化后的中间码，首先会被转为汇编代码(Plan9)，然后调用汇编器，汇编器会根据执行编译时设置的架构，调用对应代码来生成目标机器码

后续文章：

- [解析器眼中的 Go 语言](https://mp.weixin.qq.com/s/im8V1swu1Uvyzz7tKIGVWA)
  - Go 语言的词法和语法分析
- [Golang 如何进行类型检查](https://mp.weixin.qq.com/s/AetLwayJc3ncmqLX92uaGA)
- [详解 Golang 中间代码生成](https://mp.weixin.qq.com/s/caf3oVFzKT8VP2Ak81tpNg)
- [指令集架构、机器码与 Go 语言](https://mp.weixin.qq.com/s/I7764VnuUOxk8PCVrY7UeQ)

# 汇编

后续看：

[《Go 语言高级编程》](https://books.studygolang.com/advanced-go-programming-book/ch3-asm/ch3-01-basic.html)

[Go 夜读-go plan9 汇编入门](https://github.com/cch123/asmshare/blob/master/layout.md)

[A Quick Guide to Go's Assembler](https://golang.org/doc/asm)

[小米大佬讲解 Go 之运行与 Plan9 汇编](https://mp.weixin.qq.com/s/WLISnJ1J7_iKlwCed0uxYQ)
