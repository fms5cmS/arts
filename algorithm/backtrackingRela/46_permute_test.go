package backtrackingRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1 <= nums.length <= 6，输入元素有限，使用数组来标记某个索引处元素是否已使用
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			// 注意：这里不能放在上面的 for 里面！
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtracking()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtracking()
	return result
}

func TestPermute(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{input: []int{1}, output: [][]int{{1}}},
		{input: []int{0, 1}, output: [][]int{{0, 1}, {1, 0}}},
		{input: []int{1, 2, 3}, output: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.output, permute(test.input))
	}
}
