package stringRela

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// 如何找到最小重复字串？
// 当一个字符串由重复子串组成，最长相等前后缀不包含的子串就是最小重复子串
func repeatedSubstringPattern(s string) bool {
	next := getNext(s)
	// 打印前缀表
	fmt.Printf("s: %v\nt: %v\n\n", strings.Split(s, ""), next)
	length := len(s)
	if next[length-1] != 0 && length%(length-next[length-1]) == 0 {
		return true
	}
	return false
}

// s = "abc", ss = "abcabc"，掐头去尾后得到 "bcab" 判断是否包含 s，不包含则不是由重复字串组成的
// s = "abcabc", ss = "abcabcabcabc"，掐头去尾后得到 "bcabcabcab" 包含了 s，所以是由重复字串构成的
func repeatedSubstringPattern2(s string) bool {
	ss := s + s
	// 这里其实可以使用 28 中 KMP 算法来判断是否包含子串
	return strings.Contains(ss[1:len(ss)-1], s)
}

func TestRepeatedSubstringPattern(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			s:    "abab",
			want: true,
		},
		{
			s:    "aba",
			want: false,
		},
		{
			s:    "abcabcabcabc",
			want: true,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, repeatedSubstringPattern(test.s), repeatedSubstringPattern2(test.s))
	}
}
