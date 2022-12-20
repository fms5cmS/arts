package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 不同于 40、90，本题求递增子序列，是不能对原数组排序的，所以去重操作也不同！
func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		if len(path) > 1 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			//	注意：这里不能 return
		}
		// 用于标记本层某个元素是否已使用，用于去重
		set := make(map[int]struct{})
		for i := startIndex; i < len(nums); i++ {
			// 保证递增，相同元素也视为递增
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			if _, exists := set[nums[i]]; exists {
				continue
			}
			set[nums[i]] = struct{}{}
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return result
}

func TestFindSubsequences(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{input: []int{4, 6, 7, 7}, output: [][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}}},
		{input: []int{4, 4, 3, 2, 1}, output: [][]int{{4, 4}}},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.output, findSubsequences(test.input))
	}
}

// 由于题目中 -100 <= nums[i] <= 100，数量有限，所以可利用数组去重，类似 40
func findSubsequences2(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		if len(path) > 1 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			//	注意：这里不能 return
		}
		used := make([]bool, 201)
		for i := startIndex; i < len(nums); i++ {
			// 保证递增，相同元素也视为递增
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			if used[nums[i]+100] {
				continue
			}
			// 注意这里要校正！
			used[nums[i]+100] = true
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return result
}
