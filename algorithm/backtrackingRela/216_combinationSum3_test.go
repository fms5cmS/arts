package backtrackingRela

// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1-9 的正整数，并且**每种组合中不存在重复的数字**。
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0, k)
	var backtracking func(sum, startValue int)
	backtracking = func(sum, startValue int) {
		// 剪枝操作
		if sum > n {
			return
		}
		if sum == n && len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := startValue; i < 9; i++ {
			sum += i
			path = append(path, i)
			backtracking(sum, i+1)
			// 回溯
			sum -= i
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 1)
	return result
}
