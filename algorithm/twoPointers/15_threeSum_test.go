package twoPointers

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

//	给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//	注意：答案中不可以包含重复的三元组。
//
// 对数据排序后，并**保证每次移动数据时相邻两个数据的值不等，即可保证最后得到的结果不重复!
// 固定一个数，剩下两个使用双指针来查找
// 由于数据排序了，假设得到 a+b+c=0
// 那么在第一重循环 a 不变的情况下，b 增大，则 c 必然减小，故 b、c 的指针是相向移动的
// 双指针法
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	// 由于是返回值而非索引，所以可以对数组进行排序简化计算
	sort.Ints(nums)
	for i, num := range nums {
		// 最小值都大于 0 的话，后面一定没有解了
		if num > 0 {
			return result
		}
		// 对 a 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			a, b, c := nums[i], nums[left], nums[right]
			sum := a + b + c
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{a, b, c})
				// 对 b、c 做去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// 上面去重后，left 和 right 都指向了最后一个重复的元素，所以这里还需要再移动一步
				left, right = left+1, right-1
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
