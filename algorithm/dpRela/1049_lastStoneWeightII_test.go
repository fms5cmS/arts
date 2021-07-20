package dpRela

// 类似于 416
// 1 <= stones.length <= 30
// 1 <= stones[i] <= 100
func lastStoneWeightII(stones []int) int {
	// dp 数组，容量为 j 的背包最多可以装 dp[j] 这么重的石头
	// 根据 stones 的条件，最大重量为 30*100，而要求的 target 为最大重量的一半，所以这里切片长度为 1500 即可
	dp := make([]int, 1501)
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum >> 1
	for _, stone := range stones {
		for j := target; j >= stone; j-- {
			dp[j] = maxOf2Ints(dp[j], dp[j-stone]+stone)
		}
	}
	return sum - 2*dp[target]
}
