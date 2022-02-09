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
	for i, a := range nums {
		// nums 已按升序排列，如果 a > 0，三数之和必大于 0，所以退出
		if a > 0 {
			break
		}
		left, right := i+1, len(nums)-1
		if i == 0 || nums[i] != nums[i-1] { // 对 a 去重！
			for left < right {
				b, c := nums[left], nums[right]
				if a+b+c > 0 {
					right--
				} else if a+b+c < 0 {
					left++
				} else {
					// 符合条件的添加进结果中
					result = append(result, []int{a, b, c})
					// 对 b 和 c 各自去重
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					left++
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					right--
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
