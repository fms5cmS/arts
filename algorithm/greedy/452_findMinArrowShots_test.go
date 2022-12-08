package greedy

import (
	"arts/algorithm/utils"
	"sort"
)

// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0452.%E7%94%A8%E6%9C%80%E5%B0%91%E6%95%B0%E9%87%8F%E7%9A%84%E7%AE%AD%E5%BC%95%E7%88%86%E6%B0%94%E7%90%83.md
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	// 只要有气球 len(points) > 0，最少也需要一支箭
	result := 1
	for i := 1; i < len(points); i++ {
		// 相邻两个气球不重叠，需要多一支箭
		if points[i][0] > points[i-1][1] {
			result++
		} else {
			// 相邻两个气球重叠，更新重叠气球的最小右边界
			points[i][1] = utils.Min(points[i][1], points[i-1][1])
		}
	}
	return result
}
