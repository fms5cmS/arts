package backtracking

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// 注意一下结果的去重！
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(sum, startIndex int)
	backtracking = func(sum, startIndex int) {
		if sum > target {
			return
		}
		if sum == target {
			fmt.Println(path)
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		// 这里还是需要 startIndex 的，如果没有入参 startIndex，每次从 0 开始，那么返回结果会包含重复的组合
		// [2, 2, 3] 和 [3, 2, 2] 是重复的！
		for i := startIndex; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			backtracking(sum, i) // i 不用加一了，表示当前索引的数字可以重复取
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 0)
	return result
}

// 包含剪枝
func combinationSumContainsPruning(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	// 排序
	sort.Ints(candidates)
	var backtracking func(sum, startIndex int)
	backtracking = func(sum, startIndex int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		// 剪枝，排序以后，如果当前的和（sum+candidates[i]）已经超过 target 了，就不用再继续了
		for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			backtracking(sum, i) // i 不用加一了，表示当前索引的数字可以重复取
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 0)
	return result
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "2",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "3",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, combinationSum(test.candidates, test.target))
			assert.Equal(t, test.want, combinationSumContainsPruning(test.candidates, test.target))
		})
	}
}
