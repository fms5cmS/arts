package binarySearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	// case: target 的值在数组两侧
	if target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	// 两个函数最终返回值可以假设一个数组来思考！！！！
	first := getFirstPosition(nums, target)
	last := getLastPosition(nums, target)
	// case: target 的值在数组最大最小值范围内，但 target 并不在数组中
	if first > last {
		return []int{-1, -1}
	}
	// case: target 在数组中
	return []int{first, last}
}

func getFirstPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] >= target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func getLastPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] <= target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return right
}

func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// case: target 比最小值小，或 target 比最大值大
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
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

func TestRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		result []int
	}{
		{name: "1", nums: []int{5, 7, 7, 8, 8, 10}, target: 8, result: []int{3, 4}},
		{name: "2", nums: []int{5, 7, 7, 8, 8, 10}, target: 6, result: []int{-1, -1}},
		{name: "3", nums: []int{}, target: 8, result: []int{-1, -1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, searchRange(test.nums, test.target))

			assert.Equal(t, test.result, searchRange2(test.nums, test.target))
		})
	}
}
