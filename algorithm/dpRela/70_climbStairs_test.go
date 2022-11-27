package dpRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func climbStairsMoreSpace(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		if i <= 2 {
			dp[i] = i
		}
	}
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 对上面的方法可以优化空间复杂度
func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	// 爬到第 i 层楼梯，有 dp[i] 种方法
	dp := [3]int{0, 1, 2}
	for i := 3; i <= n; i++ {
		dp[1], dp[2] = dp[2], dp[1]+dp[2]
	}
	return dp[2]
}

// 和 509 几乎相同
// 爬到楼顶的方法 = 爬到倒数第一层楼梯的方法 + 爬到倒数第二层楼梯的方法！
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	sum := 0
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		sum = a + b
		a, b = b, sum
	}
	return sum
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "1",
			n:    1,
			want: 1,
		},
		{
			name: "2",
			n:    2,
			want: 2,
		},
		{
			name: "3",
			n:    3,
			want: 3,
		},
		{
			name: "4",
			n:    4,
			want: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, climbStairs(test.n))
			assert.Equal(t, test.want, climbStairs2(test.n))
			assert.Equal(t, test.want, climbStairsMoreSpace(test.n))
		})
	}
}

// 拓展：[有多少种不同的爬楼梯方法？](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0070.%E7%88%AC%E6%A5%BC%E6%A2%AF%E5%AE%8C%E5%85%A8%E8%83%8C%E5%8C%85%E7%89%88%E6%9C%AC.md)
