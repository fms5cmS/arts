package dp

import "math"

func numSquares(n int) int {
	// 和为i的完全平方数的最少数量为dp[i]
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	// 遍历顺序无所谓，这里先遍历背包再遍历物品
	for i := 0; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = minOf2Int(dp[i-j*j]+1, dp[i])
		}
	}
	return dp[n]
}
