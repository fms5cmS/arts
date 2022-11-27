package dpRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0062.%E4%B8%8D%E5%90%8C%E8%B7%AF%E5%BE%84.md

func uniquePaths(m int, n int) int {
	// dp 数组，dp[i][j] 代表从起点 (0, 0) 出发到达 (i, j) 一共有 dp[i][j] 条不同的路径
	dp := make([][]int, m)
	for row := 0; row < m; row++ {
		dp[row] = make([]int, n)
	}
	// 初始化 dp 数组，robot 只能向右、向下移动
	// 第一列 robot 到达的方式只能是向下移动，所以都是 1
	for row := 0; row < m; row++ {
		dp[row][0] = 1
	}
	// 第一行 robot 到达的方式只能是向右移动，所以都是 1
	for col := 0; col < n; col++ {
		dp[0][col] = 1
	}
	// 递推公式
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			// robot 到达非第一列且非第一行位置的方式只能是从该位置上方、该位置左侧到达
			dp[row][col] = dp[row-1][col] + dp[row][col-1]
		}
	}
	return dp[m-1][n-1]
}

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
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "1",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "2",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "3",
			m:    23,
			n:    12,
			want: 193536720,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, uniquePaths(test.m, test.n))
			assert.Equal(t, test.want, uniquePathsByDFS(test.m, test.n))
		})
	}
}
