package twoPointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	k := len(nums) - 1
	// 双指针法，结果数组 result 的最大值在原数组 nums 的两端，要么最左，要么最右
	// 注：i <= j ！因为最后要处理两个元素
	for i, j := 0, len(nums)-1; i <= j; k-- {
		if nums[i]*nums[i] < nums[j]*nums[j] {
			result[k] = nums[j] * nums[j]
			j--
		} else {
			result[k] = nums[i] * nums[i]
			i++
		}
	}
	return result
}

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{-4, -1, 0, 3, 10},
			output: []int{0, 1, 9, 16, 100},
		},
		{
			input:  []int{-7, -3, 2, 3, 11},
			output: []int{4, 9, 9, 49, 121},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, sortedSquares(test.input))
	}
}
