package examples

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"testing"
)

// New、Errorf 自定义原始错误时，会保存堆栈信息，而 Go 原生的 errors.New 不会保存堆栈信息，通常用于自己的应用代码中主动生成错误
// Wrap、Wrapf 包装原始错误，且会保存堆栈信息，通常用于和其他第三方库或标准库协作时包装他们返回的原始错误
// WithMessage 对原始错误携带一个描述信息，不再保存堆栈信息了
// Cause 获取原始错误，再使用 Is 或 As 来和 sentinel error（预定义的 error，如 io.EOF） 进行判定
// 直接返回错误，而不是每个错误产生的地方到处打日志，然后再在程序的顶部或者是工作的 goroutine 顶部（请求入口），使用 %+v 把堆栈详情打印或记录
// 如果自己的项目代码类似第三方库会有很多人使用，那么自己的项目代码不应该使用 Wrap，而应该直接返回原始错误，否则如果调用方也使用了 Wrap 来包装你返回的错误，最后打印时会有两次堆栈信息，所以仅在业务代码中使用，第三方库不能使用

func TestWrapError(t *testing.T) {
	_, err := readConfig()
	if err != nil {
		fmt.Printf("original error: %T -> %v\n", errors.Cause(err), errors.Cause(err))
		// 也可以打印 err.Error()
		fmt.Printf("withMessage : %s\n",err)
		// 注意：这里使用 %+v 的方式来打印堆栈信息
		fmt.Printf("stack trace: \n%+v\n", err)
		os.Exit(1)
	}
}

func readConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := readFile(filepath.Join(home, ".settings/xml"))
	return config, errors.WithMessage(err, "could not read config")
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer file.Close()
	buf := make([]byte, 0)
	_, err = file.Read(buf)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}
