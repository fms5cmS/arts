package twoPointer

import (
	"fmt"
	"sort"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := range nums {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			// 去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for right > left {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for right > left && nums[right] == nums[right-1] {
						right--
					}
					for right > left && nums[left] == nums[left+1] {
						left++
					}
					right--
					left++
				}
			}
		}
	}
	return result
}

func TestFourSum(t *testing.T) {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
