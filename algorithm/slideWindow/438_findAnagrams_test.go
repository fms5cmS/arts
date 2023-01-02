package slideWindow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	// 统计需要匹配的每个字符的数量
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	// 窗口的左右边界
	left, right := 0, 0
	// 窗口中每个字符的计数
	window := make(map[byte]int)
	// 窗口内验证后有多少个正确的字符，如果值和 len(need) 相等，就说明窗口内字符满足条件
	valid := 0
	for ; right < len(s); right++ {
		curChar := s[right]
		// 判断该字符是否是 p 中的字符
		if _, exists := need[curChar]; exists {
			window[curChar]++                     // 对窗口内该字符的出现数量进行统计
			if window[curChar] == need[curChar] { // 判断该字符的出现次数是否已满足条件
				valid++
			}
		}
		// 窗口已满
		for right-left+1 == len(p) {
			// 满足条件，添加到返回结果中
			if valid == len(need) {
				result = append(result, left)
			}
			// 窗口左侧元素弹出，并移动边界
			leftChar := s[left]
			left++
			// 如果左侧元素是 p 中出现过的字符，需要对 window、valid 进行调整
			if _, exists := need[leftChar]; exists {
				if window[leftChar] == need[leftChar] {
					valid--
				}
				window[leftChar]--
			}
		}
	}
	return result
}

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s, p string
		want []int
	}{
		{
			name: "1",
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},
		{
			name: "2",
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, findAnagrams(test.s, test.p))
		})
	}
}
