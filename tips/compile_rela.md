# 如何向二进制文件中加入编译时间戳、go 的版本信息？

以 Linux 系统下向编译的二进制文件中添加为例，

```go
var (
  // 这三个参数从外部在源码编译时传入
	buildstamp   = "unknown"
	gitstatus    = "unknown"
	gitcommitlog = "unknown"
	// 这里使用 runtime 包的函数来初始化值，也可以类似 buildstamp 从外部向其传值
	goversion = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("UTC Build Time : %s\n", buildstamp)
		fmt.Printf("Git Commit Hash: %s\n", gitstatus)
		fmt.Printf("Git Commit Hash: %s\n", gitcommitlog)
		fmt.Printf("Golang Version : %s\n", goversion)
		return
	}
}

```

这里使用一个 shell 脚本来完成程序内部参数赋值及可执行文件的生成：

```shell
#!/bin/bash

# 获取源码最近一次 git commit log，包含 commit sha 值，以及 commit message
gitcommitlog=`git log --pretty=oneline -n 1`
# 检查源码在git commit 基础上，是否有本地修改，且未提交的内容
gitstatus=`git status -s`
# 获取当前时间
buildstamp=`date +'%Y.%m.%d.%H%M%S'`

flags="-X 'main.buildstamp=${buildstamp}' -X 'main.gitstatus=${gitstatus}' -X 'main.gitcommitlog=${gitcommitlog}'"
# 会生成名为 add_info_to_binaryfile 的可执行文件
go build -ldflags "${flags}" -x -o add_info_to_binaryfile main.go
```

执行：`./add_info_to_binaryfile -v` 就可以看到输出了。

详见：[二进制文件加入 Git 版本的坑](https://mp.weixin.qq.com/s/VZXQeEeNNTLJPfS8tJVJZg)
