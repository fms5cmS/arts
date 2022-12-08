package dpRela

import (
	"arts/algorithm/utils"
	"strings"
)

// 01 背包，物品的重量有两个维度
// 字符串中 0、1 的数量相当于物品重量，字符串的个数相当于物品价值
func findMaxForm(strs []string, m int, n int) int {
	// dp 数组，最多有 i 个 0 和 j 个 1 的 strs 的最大子集的大小为 dp[i][j]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs { // 遍历物品
		zeroNum, oneNum := 0, 0
		zeroNum = strings.Count(str, "0")
		oneNum = strings.Count(str, "1")
		// for _, s := range str {
		//	if s == '0' {
		//		zeroNum++
		//	} else {
		//		oneNum++
		//	}
		// }
		for i := m; i >= zeroNum; i-- { // 遍历背包容量
			for j := n; j >= oneNum; j-- {
				dp[i][j] = utils.Max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}
