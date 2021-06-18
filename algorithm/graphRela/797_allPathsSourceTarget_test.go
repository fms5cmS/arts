package graphRela

// 无环图，所以不用标记某个节点是否访问过
func allPathsSourceTarget(graph [][]int) [][]int {
	rs := make([][]int, 0)
	var dfs func(int, int, []int)

	dfs = func(cur int, end int,path []int){
		if cur == end {  // 终点
			cp := make([]int,len(path))
			copy(cp, path)
			rs = append(rs, append(cp,cur))
		}
		for i:=0;i<len(graph[cur]);i++ {  // 当前节点对应的分支
			dfs(graph[cur][i],end, append(path,cur)) // 记录当前，指向下一个
		}
	}

	dfs(0,len(graph)-1, []int{})  // 当前节点
	return rs
}