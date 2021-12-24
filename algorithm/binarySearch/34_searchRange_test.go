package binarySearch

import (
	"fmt"
	"testing"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	// 找第一个位置
	left1, right1 := 0, len(nums)
	for left1 < right1 {
		mid := left1 + (right1-left1)>>1
		if nums[mid] >= target {
			right1 = mid
		} else {
			left1 = mid + 1
		}
	}
	first := right1
	// 找最后一个位置
	left2, right2 := 0, len(nums)-1
	for left2 <= right2 {
		mid := left2 + (right2-left2)>>1
		if nums[mid] <= target {
			left2 = mid + 1
		} else {
			right2 = mid - 1
		}
	}
	last := right2
	if first > last {
		return []int{-1, -1}
	}
	return []int{first, last}
}

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 8, 10}, 8))
	//fmt.Println(searchRange2([]int{1}, 1))
}

func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	first := getFirst2(nums, target)
	last := getLast2(nums, target)
	if first > last {
		return []int{-1, -1}
	}
	return []int{first, last}
}

// 这里可以参考 35
func getFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}

func getFirst2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func getLast(nums []int, target int) int {
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

func getLast2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}

func TestGetFirst(t *testing.T) {
	fmt.Println(getFirst([]int{5, 7, 7, 8, 8, 8, 10}, 8))
	fmt.Println(getFirst([]int{1}, 1))
	fmt.Println(getLast([]int{5, 7, 7, 8, 8, 8, 10}, 8))
	fmt.Println(getLast([]int{1}, 1))
}
