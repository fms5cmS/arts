package dpRela

import (
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
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 1; i < len(nums); i++ {
		maxF[i] = max(maxF[i-1]*nums[i], max(nums[i], minF[i-1]*nums[i]))
		minF[i] = min(minF[i-1]*nums[i], min(nums[i], maxF[i-1]*nums[i]))
	}
	answer := maxF[0]
	for i := 1; i < len(nums); i++ {
		answer = max(answer, maxF[i])
	}
	return answer
}

func maxProductSimple(nums []int) int {
	maxF, minF := nums[0], nums[0]
	answer := nums[0]
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF // 这里必须使用临时变量
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		answer = max(maxF, answer)
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
