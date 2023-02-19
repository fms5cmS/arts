package binarySearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 要在数组中插入目标值，有四种情况
// 目标值在数组所有元素之前
// 目标值等于数组中某一个元素
// 目标值插入数组中的位置
// 目标值在数组所有元素之后
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			// 目标值等于数组中某一个元素  return middle
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	// 此时 right < left
	// 目标值在数组所有元素之前  [0, -1]，left 始终为0，而 right 此时等于 -1
	// 目标值在数组所有元素之后的情况 [left, right]，right 始终为 len(nums)-1
	// 目标值插入数组中的位置，假设 nums = [1, 3, 5, 6], target = 2
	// 		for 循环走到
	//			left = 0, right = 1 时，此时 mid = 0, 对应的值为 1 ( < target)，于是 left 更新为 1
	// 			left = 1, right = 1 时，此时 mid = 1, 对应的值为 3 ( > target)，于是 right 更新为 0
	//			left = 1, right = 0 时，退出循环
	//			要插入的位置应为 1
	return right + 1 // 这里可以直接返回 left
}

func searchInsert1(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	// 注意，与上面不同，这里直接返回 right 即可！
	return right
}

// 暴力解法
func searchInsert2(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "1",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "2",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "3",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			name:   "4",
			nums:   []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, searchInsert(test.nums, test.target))
			assert.Equal(t, test.want, searchInsert1(test.nums, test.target))
		})
	}
}
