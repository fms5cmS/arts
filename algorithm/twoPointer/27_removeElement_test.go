package twoPointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if val != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums   []int
		val    int
		output int
	}{
		{
			nums:   []int{3, 2, 2, 3},
			val:    3,
			output: 2,
		},
		{
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			output: 5,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, removeElement(test.nums, test.val))
	}
}
