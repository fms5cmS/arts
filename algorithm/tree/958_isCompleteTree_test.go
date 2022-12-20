package tree

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur != nil {
			queue = append(queue, cur.Left, cur.Right)
			continue
		}
		// cur == nil，如果后面存在不为 nil 的节点，则代表不是完全二叉树
		if len(queue) > 0 {
			cur2 := queue[0]
			if cur2 != nil {
				result = false
			}
		}
		if !result {
			return result
		}
	}
	return result
}
