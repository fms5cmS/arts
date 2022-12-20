package string

import (
	"fmt"
	"strings"
	"testing"
)

// 把字符串 s 中的每个空格替换成"%20"。
func replaceSpace(s string) string {
	spaceCount := strings.Count(s, " ")
	str := make([]byte, 0, len(s)+spaceCount*2)
	for _, char := range s {
		if char == ' ' {
			str = append(str, []byte("%20")...)
		} else {
			str = append(str, byte(char))
		}
	}
	return string(str)
}

// C++ 中字符串支持修改，所以可以以 O(1) 的空间复杂度实现，不过在使用双指针时需要从后往前遍历，
// 如果是从前往后遍历，因为每次修改空格的字符时需要将后面的字符向后移动，所以复杂度会是 O(n^2)
// Go 中字符串不支持修改，所以新建一个 []byte 后可以从前往后的填充值
func replaceSpaceTwoPointer(s string) string {
	spaceCount := strings.Count(s, " ")
	str := make([]byte, len(s)+spaceCount*2)
	for i, j := len(str)-1, len(s)-1; j >= 0; {
		if s[j] != ' ' {
			str[i] = s[j]
		} else {
			str[i], str[i-1], str[i-2] = '0', '2', '%'
			i -= 2
		}
		i--
		j--
	}
	return string(str)
}

func TestReplaceSpace(t *testing.T) {
	fmt.Println(replaceSpaceTwoPointer("We are happy."))
}
