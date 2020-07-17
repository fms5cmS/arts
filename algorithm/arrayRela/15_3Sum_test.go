package arrayRela

import (
	"sort"
	"testing"
)

// 15. 3Sum
// 对数据排序后，并保证每次移动数据时相邻两个数据的值不等，即可保证最后得到的结果不重复!
// 固定一个数，剩下两个使用双指针来查找
// 由于数据排序了，假设得到 a+b+c=0
// 那么在第一重循环 a 不变的情况下，b 增大，则 c 必然减小，故 b、c 的指针是相向移动的
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		target := 0 - nums[i]
		if nums[i] > 0 {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			for l < r {
				if nums[l]+nums[r] == target {
					result = append(result, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if nums[l]+nums[r] < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return result
}

func threeSum_old(nums []int) [][]int {
	result := make([][]int, 0)
	length := len(nums)
	sort.Ints(nums)
	for first := 0; first < length; first++ {
		// 要保证同一指针相邻两次得到的值不等
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < length; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			third := length - 1
			for second < third && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}
			if second == third {
				break
			}
			if nums[first]+nums[second]+nums[third] == 0 {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return result
}

func TestThreeSum(t *testing.T) {
	nums := []int{1, -1, -1, 0}
	ret := threeSum(nums)
	t.Log(ret)
}
