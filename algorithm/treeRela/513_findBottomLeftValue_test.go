package treeRela

func findBottomLeftValue(root *TreeNode) int {
	queue, result := make([]*TreeNode, 0), 0
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 记录最后一行第一个元素
			if i == 0 {
				result = cur.Val
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return result
}
