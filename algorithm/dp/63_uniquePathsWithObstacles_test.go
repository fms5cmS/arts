package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 1. dp 数组，dp[i][j] 代表从起点 (0, 0) 出发到达 (i, j) 一共有 dp[i][j] 条不同的路径
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}
	// 3. 初始化，这里在条件中增加了有无障碍的判断，如果某个位置有了障碍物，同一行内，该位置后面的 dp[i][0] 都应该是 0，因为走不过去了!!!
	for row := 0; row < len(obstacleGrid) && obstacleGrid[row][0] == 0; row++ {
		dp[row][0] = 1
	}
	for col := 0; col < len(obstacleGrid[0]) && obstacleGrid[0][col] == 0; col++ {
		dp[0][col] = 1
	}
	// 4. 遍历顺序
	// 注意：这里都是从 1 开始的
	for row := 1; row < len(dp); row++ {
		for col := 1; col < len(dp[row]); col++ {
			// 2. 递推公式，只有无障碍时才更新，否则其位置的值保持为默认值 0
			if obstacleGrid[row][col] == 1 {
				continue
			}
			dp[row][col] = dp[row-1][col] + dp[row][col-1]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name         string
		obstacleGrid [][]int
		want         int
	}{
		{
			name:         "1",
			obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want:         2,
		},
		{
			name:         "2",
			obstacleGrid: [][]int{{0, 1}, {0, 0}},
			want:         1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, uniquePathsWithObstacles(test.obstacleGrid))
		})
	}
}
