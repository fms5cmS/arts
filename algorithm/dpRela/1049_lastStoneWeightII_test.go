package dpRela

import (
	"arts/algorithm/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 类似于 416
// eg：[2, 7, 4, 1, 8, 1]
// 1. 组合 2 和 4，得到 4-2=2
// 2. 组合 7 和 8，得到 8-7=1
// 3. 组合上面得到的 2 和 1，得到 2-1=1
// 4. 组合原数组中的 1 和 1，得到 1-1=0
// 5. 仅剩第四步剩下的 1
// 所有公式组合起来就是 ((4-2)-(8-7))+(1-1) = 1
// 解开上面的括号，得到 -2 + 7 +4 + 1 - 8 - 1 = 7+4+1-1-2-8 = (7+4+1)-(1+2+8)
// 所以题目就是要将所有的数字分成两组，求两组数字和的最小差
// 可以判断出，**如果其中一组数字越接近所有数字之和，那么此时两组数字之差就越小！**
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum >> 1
	// dp 数组代表容量为 i 的背包里最多可以放 dp[i] 重量的物品
	dp := make([]int, target+1)
	for _, stone := range stones {
		for i := target; i >= stone; i-- {
			dp[i] = utils.Max(dp[i], dp[i-stone]+stone)
		}
	}
	return sum - 2*dp[target]
}

func TestLastStoneWeightII(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{
			name:   "1",
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
		{
			name:   "2",
			stones: []int{31, 26, 33, 21, 40},
			want:   5,
		},
		{
			name:   "3",
			stones: []int{1, 2, 1, 7, 9, 4},
			want:   0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, lastStoneWeightII(test.stones))
		})
	}
}
