package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		// 添加操作要放在终止条件之前，否则会漏掉本身
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		// 终止条件
		if startIndex >= len(nums) {
			return
		}
		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return result
}

func TestSubSet(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output [][]int
	}{
		{name: "1", input: []int{1, 2, 3}, output: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}},
		{name: "2", input: []int{0}, output: [][]int{{}, {0}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, subsets(test.input))
		})
	}
}
