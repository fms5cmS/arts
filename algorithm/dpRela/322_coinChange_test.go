package dpRela

import "math"

func coinChange(coins []int, amount int) int {
	// 1. 递推公式：dp[i] 代表凑足金额为 i 所需钱币的最少个数为 dp[i]
	dp := make([]int, amount+1)
	// 3. 初始化数组，dp[0] = 0，其他的值初始化为最大整型（下面递推公式需要取得最小值，这里赋值为最大整型防止被初始值覆盖）
	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt64
	}
	// 遍历钱币
	for _, coin := range coins {
		// 遍历金额
		for i := coin; i <= amount; i++ {
			if dp[i-coin] != math.MaxInt64 {
				// 递推公式，因为 dp[i] 的含义（见上），所以这里需要取最小值
				dp[i] = minOf2Int(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func minOf2Int(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 这个题目要求的是最小个数，钱币有序或无序并不影响最小个数，所以也可以先遍历金额，再遍历钱币！
// 学习自：https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0322.%E9%9B%B6%E9%92%B1%E5%85%91%E6%8D%A2.md