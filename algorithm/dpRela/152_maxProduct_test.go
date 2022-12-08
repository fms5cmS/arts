package dpRela

import (
	"arts/algorithm/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxProduct(nums []int) int {
	// 需要两个数组，一个记录之前最大乘积，一个记录之前最小乘积
	// 其中最小乘积是为了在当前位置上值为负数，且之前的最小乘积也为负数时也可能得到乘积最大
	// 可以结合下面第一个测试用例理解
	maxF := make([]int, len(nums))
	minF := make([]int, len(nums))
	copy(maxF, nums)
	copy(minF, nums)
	for i := 1; i < len(nums); i++ {
		maxF[i] = utils.Max(maxF[i-1]*nums[i], utils.Max(nums[i], minF[i-1]*nums[i]))
		minF[i] = utils.Min(minF[i-1]*nums[i], utils.Min(nums[i], maxF[i-1]*nums[i]))
	}
	answer := maxF[0]
	for i := 1; i < len(nums); i++ {
		answer = utils.Max(answer, maxF[i])
	}
	return answer
}

func maxProductSimple(nums []int) int {
	maxF, minF := nums[0], nums[0]
	answer := nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF // 这里必须使用临时变量
		maxF = utils.Max(mx*nums[i], utils.Max(nums[i], mn*nums[i]))
		minF = utils.Min(mn*nums[i], utils.Min(nums[i], mx*nums[i]))
		answer = utils.Max(maxF, answer)
	}
	return answer
}

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "1",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "2",
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "3",
			nums: []int{-4, -3, -2},
			want: 12,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, maxProduct(test.nums))
			assert.Equal(t, test.want, maxProductSimple(test.nums))
		})
	}
}
