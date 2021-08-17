package binarySearch

import (
	"fmt"
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
	// 目标值插入数组中的位置 [left, right]，return  right + 1
	// 目标值在数组所有元素之后的情况 [left, right]，right 始终为 len(nums)-1
	return right + 1
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
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 4))
}
