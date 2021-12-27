package backtrackingRela

import (
	"fmt"
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
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
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

// 注意，这个示例的返回为 [2,2,3],[7]，而 [2,3,2] 是重复的
func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	result := combinationSum(candidates, 7)
	fmt.Println(result)
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
