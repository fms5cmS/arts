package dp

import (
	"arts/algorithm/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	// dp[start][end] 代表 s 中从 start 到 end 最长的回文子序列长度
	// 动态规划的计算方向，根据计算方向来决定遍历顺序！！
	// 最终的结果保存在 dp[0][len(s)-1] 上
	// 而 dp[0][len(s)-1] 的值是根据 dp[start][end-1]、dp[start+1][end-1]、dp[start+1][end] 推导出来的
	for start := len(s) - 1; start >= 0; start-- {
		for end := start + 1; end < len(s); end++ {
			if s[start] == s[end] {
				dp[start][end] = 2 + dp[start+1][end-1] // 相比原来多了两个字符
			} else {
				dp[start][end] = utils.Max(dp[start+1][end], dp[start][end-1])
			}
			fmt.Printf("(%d, %d) = %d, str = %s\n", start, end, dp[start][end], s[start:end+1])
		}
	}
	for _, vs := range dp {
		fmt.Println(vs)
	}
	return dp[0][len(s)-1]
}

func TestLongestPalindromeSubseq(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "1",
			s:    "bbbab",
			want: 4,
		},
		{
			name: "2",
			s:    "cbbd",
			want: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, longestPalindromeSubseq(test.s))
		})
	}
}
