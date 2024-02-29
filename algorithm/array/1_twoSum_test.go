package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 由于返回的是数组的索引，所以不能对数组排序后使用双指针方式
func twoSum(nums []int, target int) []int {
	// 存储已经遍历过的数据！
	recordMap := make(map[int]int)
	// 注意里面的顺序！！！存储 num 和 i 的逻辑要放在后面，可以参考 [3, 3] target=6 来理解
	for i, num := range nums {
		if index, exists := recordMap[target-num]; exists {
			return []int{index, i}
		}
		recordMap[num] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, twoSum(test.nums, test.target))
		assert.Equal(t, test.output, twoSumForce(test.nums, test.target))
	}
}

// 暴力解法 O(n^2)
func twoSumForce(nums []int, target int) []int {
	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			if a+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
