package dpRela

import (
	"arts/algorithm/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 找是否可以将这个数组分割成两个子集，使得两个子集的元素和相等
// dp 数组，容量为 i 的背包，放入的子集之和的最大值为 dp[i]
// dp[i] == i 判断容量为 i 的背包是否装满，如 nums = [1, 5, 11, 5]，dp[7] 只能装 6（放入了 1 和 5），此时背包就是未装满
func canPartition(nums []int) bool {
	sum := 0
	// 计算背包最大容量
	for _, num := range nums {
		sum += num
	}
	// 如果物品总重为奇数，则必然不可分割
	if sum%2 == 1 {
		return false
	}
	target := sum >> 1
	// dp 数组初始化
	dp := make([]int, target+1)
	for _, num := range nums { // 遍历物品
		for j := target; j >= num; j-- { // 遍历背包，倒序遍历是为了保证当前物品仅会被放入一次
			// 递推公式
			// dp[j] 代表不放入当前物品，dp[j-num]+num 代表放入当前物品
			dp[j] = utils.Max(dp[j], dp[j-num]+num)
		}
	}
	return dp[target] == target
}

func canPartition2(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, num := range nums {
		for i := sum; i > 0; i-- {
			if i >= num {
				dp[i] = dp[i] || dp[i-num]
			}
		}
	}
	return dp[sum]
}

func TestCanPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "1",
			nums: []int{1, 5, 11, 5},
			want: true,
		},
		{
			name: "2",
			nums: []int{1, 2, 3, 5},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, canPartition(test.nums))
			assert.Equal(t, test.want, canPartition2(test.nums))
		})
	}
}
