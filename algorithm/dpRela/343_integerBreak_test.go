package dpRela

import (
	"fmt"
	"testing"
)

func integerBreak(n int) int {
	// 1. dp 数组，dp[i] 代表分拆数字 i 可以得到的最大乘积为 dp[i]
	dp := make([]int, n+1)
	// n 不小于 2 且不大于 58，所以这里不会越界
	// 3. dp 数组初始化，这里仅初始化 dp[2]，从 dp 数组的定义看，dp[0]、dp[1] 是没有意义的数值，所以不必初始化
	dp[2] = 1
	// 4. 从小到大 遍历 n
	for i := 3; i <= n; i++ {
		// j < i-1 保证了拆分开的数字都是正整数，而不会有 0
		for j := 1; j < i-1; j++ {
			// 2. 递推公式
			// i 拆分为 i-j、j，其乘积为 (i-j)*j    这是拆分为两个正整数相乘
			// i 拆分为 j、(i-j) 拆分的结果，其乘积为 dp[i-j]*j    这是拆分为两个及两个以上的正整数相乘
			dp[i] = maxOf2Ints(dp[i], maxOf2Ints((i-j)*j, dp[i-j]*j))
		}
	}
	for i, v := range dp {
		if i <= 1 {
			continue
		}
		fmt.Println(i, " 拆分后最大乘积为 ", v)
	}
	return dp[n]
}

func maxOf2Ints(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestIntegerBreak(t *testing.T) {
	integerBreak(10)
}
