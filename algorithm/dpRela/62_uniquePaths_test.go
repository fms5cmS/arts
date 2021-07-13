package dpRela

import (
	"testing"
)

// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0062.%E4%B8%8D%E5%90%8C%E8%B7%AF%E5%BE%84.md
// 深度优先遍历
func uniquePathsByDFS(m int, n int) int {
	var dfs func(i, j, m, n int) int
	// i,j 代表了起点，m,n 代表了终点
	dfs = func(i, j, m, n int) int {
		// 越界
		if i > m || j > n {
			return 0
		}
		// 找到了一种方法
		if i == m && j == n {
			return 1
		}
		// 只能向下或向右移动
		return dfs(i+1, j, m, n) + dfs(i, j+1, m, n)
	}
	return dfs(1, 1, m, n)
}

func TestUniquePath(t *testing.T) {
	t.Log(uniquePathsByDFS(3, 3))
	t.Log(uniquePaths(3, 3))
}

func uniquePaths(m int, n int) int {
	// 1. dp 数组，dp[i][j] 代表从起点 (0, 0) 出发到达 (i, j) 一共有 dp[i][j] 条不同的路径
	dp := make([][]int, m)
	// 3. dp 数组初始化
	for i := range dp {
		dp[i] = make([]int, n)
		// 从 (0, 0) 到 (i, 0) 的路径只有一条，机器人之恶能一直向下走（机器人只能向下或向右两个方向移动）
		dp[i][0] = 1
		for j := range dp[i] {
			// 同 dp[i][0] 类似
			dp[0][j] = 1
		}
	}
	// 4. 遍历顺序从递推公式来看，dp[i][j] 都是从其上方和左方推导而来，那么从左到右一层一层遍历就可以了。
	// 注意：这里都是从 1 开始的
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 2. 递推公式，机器人只能向下或向右两个方向！所以 dp[i][j] 可以由两个方向上的结果推导出来
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
