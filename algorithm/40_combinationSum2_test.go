package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func combinationSum2(candidates []int, target int) [][]int {
	used := make([]bool, len(candidates))
	// 把 candidates 排序，让其相同的元素都挨在一起
	sort.Ints(candidates)
	// 记录每次符合要求的组合
	path := make([]int, 0)
	// 返回结果
	result := make([][]int, 0)
	backtracking(candidates, target, 0, 0, used, path, &result)
	return result
}

func backtracking(candidates []int, target, sum, startIndex int, used []bool, path []int, result *[][]int) {
	// 满足要求，将 path 中的组合复制后放入 result
	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	// sum 过大，不必再进行操作
	if sum > target {
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		// used[i-1] == true，说明同一树支candidates[i - 1]使用过
		// used[i-1] == false，说明同一树层candidates[i - 1]使用过
		// 要对同一树层使用过的元素进行跳过
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		path = append(path, candidates[i])
		sum += candidates[i]
		used[i] = true
		// 递归
		backtracking(candidates, target, sum, i+1, used, path, result)
		fmt.Println(path)
		// 回溯
		path = path[:len(path)-1]
		sum -= candidates[i]
		used[i] = false
	}
}

func TestCombinationSum2(t *testing.T) {

}
