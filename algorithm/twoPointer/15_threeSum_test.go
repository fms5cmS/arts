package twoPointer

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

//  给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//  注意：答案中不可以包含重复的三元组。
// 对数据排序后，并**保证每次移动数据时相邻两个数据的值不等，即可保证最后得到的结果不重复!
// 固定一个数，剩下两个使用双指针来查找
// 由于数据排序了，假设得到 a+b+c=0
// 那么在第一重循环 a 不变的情况下，b 增大，则 c 必然减小，故 b、c 的指针是相向移动的
// 双指针法
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	// 找出三数之和等于 0 的，nums[i]+nums[left]+nums[right]=0
	for i, num := range nums {
		// 如果三数之中最小的值 nums[i] > 0，则无论如何组合都不可能凑成，直接返回 result 中已有的组合
		if num > 0 {
			return result
		}
		// 去重，nums[i] 即为 num，这里是为了方便看比较的两个值
		// 这里不能用 nums[i] == nums[i+1]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for right > left {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 去重
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				for right > left && nums[left] == nums[left+1] {
					left++
				}
				// 双指针收缩
				right--
				left++
			}
		}
	}
	return result
}

// map 实现的去重很难直接写出没有 bug 的代码
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
	testCases := []struct {
		nums   []int
		result [][]int
	}{
		{
			nums:   []int{-1, 0, 1, 2, -1, -4},
			result: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:   []int{0, 1, 1},
			result: [][]int{},
		},
		{
			nums:   []int{0, 0, 0},
			result: [][]int{{0, 0, 0}},
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, len(testCase.result), len(threeSum(testCase.nums)))
	}
}
