package stackAndQueueRela

import (
	"fmt"
	"strings"
	"testing"
)

// 必须使用栈结构！每次遇到 .. 就弹出一个目录
func simplifyPath(path string) string {
	stack := make([]string, 0)
	for _, name := range strings.Split(path, "/") {
		if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if name != "" && name != "." {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func TestSimplifyPath(t *testing.T) {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/../"))
}
