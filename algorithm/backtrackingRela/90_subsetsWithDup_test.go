package backtrackingRela

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// 需要做去重操作，可以参考 40
func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	// 排序
	sort.Ints(nums)
	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		for i := startIndex; i < len(nums); i++ {
			// 去重操作
			// 假设输入参数为 [1, 2, 2]，path = [1]，startIndex = 1
			// i = 1 时，会得到 [1, 2]，此时的 2 是索引为 1 的 2
			// 当回溯时，path = [1]，i = 2，且 i > startIndex，因为已经有了 [1, 2] 所以这里直接 continue
			if i > startIndex && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return result
}

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{input: []int{1, 2, 2}, output: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{input: []int{0}, output: [][]int{{}, {0}}},
	}
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.output, subsetsWithDup(test.input))
	}
}
