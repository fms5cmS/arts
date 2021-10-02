package dpRela

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
	dp := [3]int{0, 1, 2}
	for i := 3; i <= n; i++ {
		dp[1], dp[2] = dp[2], dp[1]+dp[2]
	}
	return dp[2]
}

// 拓展：[有多少种不同的爬楼梯方法？](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0070.%E7%88%AC%E6%A5%BC%E6%A2%AF%E5%AE%8C%E5%85%A8%E8%83%8C%E5%8C%85%E7%89%88%E6%9C%AC.md)
