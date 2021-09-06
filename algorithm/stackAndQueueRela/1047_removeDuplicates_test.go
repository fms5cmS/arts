package stackAndQueueRela

import (
	"fmt"
	"testing"
)

// 这种匹配问题适合使用栈来完成
func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		// 判断栈顶元素与当前字符是否相等，相等则将栈顶元素出栈
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			// 不等，则将当前元素入栈
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates("abbaxc"))
}
