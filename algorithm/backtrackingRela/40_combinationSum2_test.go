package backtrackingRela

import (
	"fmt"
	"sort"
	"testing"
)

// 不同于 39，本题中 candidates 有重复的数字，每个元素只能使用一次
// 同样要去重
func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(candidates))
	// 排序，保证重复元素在一起
	sort.Ints(candidates)
	var backtracking func(sum, startIndex int)
	backtracking = func(sum, startIndex int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		// 假设 candidates = [1, 1, 2], target = 3
		// 先选了索引为 0 的值 1，此时 used[0] = true，在 backtracking(sum, i+1) 中，是可以继续选择索引为 1 的值 1 的，因为每个元素都可以使用一次
		//   注，这里可以得到一个结果 [1, 2]，不过这个 1 是 candidates 中索引为 0 的值 1
		// 没选索引为 0 的值 1，此时 used[0] = false，在 backtracking(sum, i+1) 中，不可以继续选择索引为 1 的值 1 了
		//   因为最后的结果需要对组合去重，而之前已经根据相同值 1（candidates 中索引为 0）处理过得到了 [1, 2] 的结果，所以不能再获取 [1, 2] 了，即使此时的 1 是 candidates 中索引为 1 的值 1
		for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
			// 要对同一树层使用过的元素进行跳过
			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}
			sum += candidates[i]
			path = append(path, candidates[i])
			used[i] = true
			backtracking(sum, i+1) // 每个元素只能使用一次，所以 i+1
			used[i] = false
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 0)
	return result
}

func TestCombinationSum2(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	result := combinationSum2(candidates, 8)
	fmt.Println(result)
}
