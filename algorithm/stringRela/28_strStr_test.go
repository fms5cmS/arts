package stringRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func strStr(haystack string, needle string) int {
	next := getNext(needle)
	// i 和 j 分别是在 haystack 和 needle 移动的指针
	i, j := 0, 0
	for ; i < len(haystack); i++ {
		// 遇到不匹配，模式串回退
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

// getNext 构造前缀表，s 为模式串
// 前缀表 next 记录了 index <= j 的字串其相同前后缀的长度
// s = "aabaaf"
//          后缀（只包含尾字符，不包含首字符的所有子串）   前缀（只包含首字符，不包含尾字符的所有子串）        最长公共前后缀   前缀表（最长公共前后缀长度）
//  a                        无                             无                                     无            0
//  aa                       a                              a                                     a             1
//  aab           		   b, ab                           a, aa                                  无            0
//  aaba                a, ba, aba                       a, aa, aab                               a             1
//  aabaa             a, aa, baa, abaa                a, aa, aab, aaba                            aa            2
//  aabaaf       f, af, aaf, baaf, abaaf            a, aa, aab, aaba, aabaa                       无            0
func getNext(s string) []int {
	next := make([]int, len(s))
	// 首字符没有前后缀，所以前缀表的长度为 0
	// next[0] = 0
	// i 指向后缀末尾，j 指向前缀末尾
	i, j := 1, 0
	for ; i < len(s); i++ {
		// 前后缀不等，j 回退
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		// 更新 next
		next[i] = j
	}
	return next
}

func TestGetNext(t *testing.T) {
	assert.Equal(t, []int{0, 1, 0, 1, 2, 0}, getNext("aabaaf"))
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			want:     -1,
		},
		{
			haystack: "a",
			needle:   "a",
			want:     0,
		},
		{
			haystack: "abc",
			needle:   "c",
			want:     2,
		},
		{
			haystack: "aaa",
			needle:   "aaaa",
			want:     -1,
		},
		{
			haystack: "mississippi",
			needle:   "issip",
			want:     4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, strStr(test.haystack, test.needle))
	}
}
