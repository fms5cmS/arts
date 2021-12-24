package twoPointer

import "sort"

//  给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//  注意：答案中不可以包含重复的三元组。
// 双指针法
func threeSumTwoPointer(nums []int) [][]int {
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
