package binarySearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	left, right := getLeft(nums, target), getRight(nums, target)
	if left > right {
		return []int{-1, -1}
	}
	return []int{left, right}
}

// 寻找左边界
func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left // right+1
}

// 寻找右边界
func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func TestRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 8, result: []int{3, 4}},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 6, result: []int{-1, -1}},
		{nums: []int{}, target: 8, result: []int{-1, -1}},
	}
	assert := assert.New(t)
	for _, test := range tests {
		actual := searchRange(test.nums, test.target)
		assert.Equal(test.result, actual)
	}
}

func TestGet(t *testing.T) {
	assert.Equal(t, 3, getLeft([]int{5, 7, 7, 8, 8, 8, 10}, 8))
	assert.Equal(t, 0, getLeft([]int{1}, 1))

	assert.Equal(t, 5, getRight([]int{5, 7, 7, 8, 8, 8, 10}, 8))
	assert.Equal(t, 0, getRight([]int{1}, 1))
}
