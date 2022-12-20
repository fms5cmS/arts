package dp

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 没有方案
	if target > sum || (target+sum)%2 == 1 {
		return 0
	}
	bagSize := (target + sum) / 2
	// dp 数组，填满容积为 j 得背包，有 dp[j] 中方案
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for _, num := range nums {
		for j := bagSize; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[bagSize]
}
