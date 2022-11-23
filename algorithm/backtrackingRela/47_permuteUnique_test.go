package backtrackingRela

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)
	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			// 这里是返回结果的去重逻辑，最后面使用 !used[i-1] 和 used[i-1] 都可以，但是不能不写。以输入 [1, 1, 2] 说明
			// !used[i-1]
			//      当 path[0] 的值是 nums[0] 时，其实已经有了 [1,1,2], [1,2,1] 的结果了
			//      所以当 path[0] 的值是 nums[1] 时，得到的结果和上面时重复的，可以跳过
			//      效率会更高
			// used[i-1]
			//      当 path[0] 的值是 nums[0] 时，只能得到 [1,2,1]
			//      当 path[0] 的值是 nums[1] 时，才能得到 [1,1,2]
			//      注意，当 path[0] 的值是 nums[2] 时，得到的 [2,1,1] 实际为 nums[2], nums[1], nums[0]
			// 如果不写，则重复元素无法放入，类似 [1,1,2] 时第二个 1 就没了，所以是错误的
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
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

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output [][]int
	}{
		{name: "1", input: []int{1, 1, 2}, output: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{name: "2", input: []int{1, 2, 3}, output: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, permuteUnique(test.input))
		})
	}
}
