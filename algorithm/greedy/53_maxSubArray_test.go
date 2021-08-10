package greedy

import "math"

// 贪心算法
func maxSubArray(nums []int) int {
	result := math.MinInt32
	count := 0
	for _, num := range nums {
		count += num
		if count > result {
			result = count
		}
		// 如果 count 计算为负数，舍弃之前的结果从下一个 num 重新开始计算，因为负数与后面的 num 计算一定会拉低结果
		if count <= 0 {
			count = 0
		}
	}
	return result
}

// 动态规划
func maxSubArrayByDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// dp[i] 表示包括 i 之前的最大连续子序列之和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxOf2Ints(dp[i-1]+nums[i], nums[i])
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func maxOf2Ints(x, y int) int {
	if x > y {
		return x
	}
	return y
}
