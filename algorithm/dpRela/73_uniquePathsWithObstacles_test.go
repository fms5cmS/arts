package dpRela

import "testing"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 1. dp 数组，dp[i][j] 代表从起点 (0, 0) 出发到达 (i, j) 一共有 dp[i][j] 条不同的路径
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}
	// 3. 初始化，这里在条件中增加了有无障碍的判断，如果某个位置有了障碍物，同一行内，该位置后面的 dp[i][0] 都应该是 0，因为走不过去了
	for i := 0; i < len(obstacleGrid) && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < len(obstacleGrid[0]) && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}
	// 4. 遍历顺序
	// 注意：这里都是从 1 开始的
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			// 2. 递推公式，只有无障碍时才更新，否则其位置的值保持为默认值 0
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func TestUniquePathsWithObstacles(t *testing.T) {
	t.Log(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	t.Log(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}
