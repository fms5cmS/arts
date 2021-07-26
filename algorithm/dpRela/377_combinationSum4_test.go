package dpRela

import "math"

func combinationSum4(nums []int, target int) int {
	// 凑成目标正整数为i的排列个数为dp[i]
	dp := make([]int, target+1)
	dp[0] = 1
	// 整数的个数不限，完全背包；得到的是排列，需要考虑整数之间的顺序，所以，先遍历背包再遍历物品
	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 && dp[i] < math.MaxInt64-dp[i-num] {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
