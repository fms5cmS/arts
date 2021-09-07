package treeRela

func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0{
		size := len(queue)
		sum := 0
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			sum += cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		result = append(result, float64(sum)/float64(size))
	}
	return result
}

