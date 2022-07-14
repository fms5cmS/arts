package arrayRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSum(nums []int, target int) []int {
	recordMap := make(map[int]int)
	for i, num := range nums {
		if index, exists := recordMap[target-num]; exists {
			return []int{index, i}
		}
		recordMap[num] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, twoSum(test.nums, test.target))
	}
}
