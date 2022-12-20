package dp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 时间和空间复杂度都是 O(n^2)
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	answer := 0
	dp := make([][]bool, len(s))
	// 实际只会使用到整个二维数组右上部分，一半的空间用不到
	for i := range dp {
		dp[i] = make([]bool, len(s))
		// 每个单独的字符都是一个回文子串
		dp[i][i] = true
		answer++
	}
	// 由于需要确定子串的位置，所以需要两个变量用于约束和确定子串，一个是起始位置 start，一个是结束位置 end
	// dp[start][end] 表示子串 [start, end] 是否为回文字符串
	for end := 1; end < len(s); end++ {
		for start := 0; start < end; start++ {
			// 新的起始位置和结束位置的字符相等时 s[end] == s[start]：
			//    如果子串长度小于 3，则现在的一定是回文子串
			//    如果其原本的字符串是回文子串 dp[start+1][end-1]，则现在的一定是回文子串
			dp[start][end] = (s[end] == s[start]) && (end-start < 3 || dp[start+1][end-1])
			if dp[start][end] {
				answer++
			}
		}
	}
	for _, rows := range dp {
		fmt.Println(rows)
	}
	return answer
}

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "1",
			s:    "abc",
			want: 3,
		},
		{
			name: "2",
			s:    "aaa",
			want: 6,
		},
		{
			name: "3",
			s:    "fdsklf",
			want: 6,
		},
		{
			name: "4",
			s:    "aaaaa",
			want: 15,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, countSubstrings(test.s))
		})
	}
}
