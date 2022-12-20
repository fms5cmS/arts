package dp

import (
	"arts/algorithm/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 题目：
// 注意：最后一个台阶并不是楼顶！初始阶梯可以从 0 或 1 开始
// 示例一：cost = [10, 15, 20]，输出 15
//       选择下标为 1 的元素作为初始阶梯，支付离开 1 阶的体力 15，跨两（！）步达到楼顶
// 示例二：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]，输出 6
//       选择下标为 0 的元素作为初始阶梯
//       1. 支付离开 0 阶的体力 1，跨两步到达 2 阶
//       2. 支付离开 2 阶的体力 1，跨两步到达 4 阶
//       3. 支付离开 4 阶的体力 1，跨两步到达 6 阶
//       4. 支付离开 6 阶的体力 1，跨一步到达 7 阶
//       5. 支付离开 7 阶的体力 1，跨两步到达 9 阶
//       6. 支付离开 9 阶的体力 1，跨一步到达楼顶
//                   总体力为 6

func minCostClimbingStairs(cost []int) int {
	// dp 数组，dp[i] 代表到达第 i 个台阶所花费的最小体力为 dp[i]
	dp := make([]int, len(cost))
	// 初始化 dp 数组
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(dp); i++ {
		// 递推公式，这里还需要加上 cost[i]，代表从 i 阶离开消耗的体力值
		dp[i] = cost[i] + utils.Min(dp[i-1], dp[i-2])
	}
	// fmt.Println(dp)
	return utils.Min(dp[len(dp)-1], dp[len(dp)-2])
}

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			name: "1",
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			name: "2",
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, minCostClimbingStairs(test.cost))
		})
	}
}
