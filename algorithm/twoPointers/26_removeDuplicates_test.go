package twoPointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow, fast := 0, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	return slow + 1
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 2},
			expected: 2,
		},
		{
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, removeDuplicates(test.nums))
	}
}
