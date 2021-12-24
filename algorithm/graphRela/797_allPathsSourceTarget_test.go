package graphRela

// 无环图，所以不用标记某个节点是否访问过
// 输入：graph = [[1,2],[3],[3],[]]
// 二维数组的 row 代表了某个顶点（vertex），row 指向的一维数组代表了该顶点指向的其他顶点
// 值为 0 的顶点指向了值为 1、2 的两个顶点
// 值为 1、2 的顶点指向了值为 3 的顶点
// 值为 3 的顶点不指向任何顶点
// 输出：[[0,1,3],[0,2,3]] 代表了所有从 0 到 3 的路径
func allPathsSourceTarget(graph [][]int) [][]int {
	rs := make([][]int, 0)
	var dfs func(int, int, []int)

	dfs = func(cur int, end int, path []int) {
		// 走到了终点
		if cur == end {
			cp := make([]int, len(path))
			copy(cp, path) // 注意这里要 copy
			rs = append(rs, append(cp, cur))
		}
		// graph[cur] 代表了当前顶点 cur 可以走到的所有其他顶点
		// 对 cur 可以走到的每个顶点 dfs
		for i := 0; i < len(graph[cur]); i++ {
			dfs(graph[cur][i], end, append(path, cur))
		}
	}

	dfs(0, len(graph)-1, []int{}) // 当前节点
	return rs
}
