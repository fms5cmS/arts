package treeRela

import "math"

func largestValues(root *TreeNode) []int {
	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		max := math.MinInt32
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Val > max {
				max = cur.Val
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		result = append(result, max)
	}
	return result
}
