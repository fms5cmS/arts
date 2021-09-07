package treeRela

// 使用迭代法的话，使用层序遍历是最为合适的，因为最大的深度就是二叉树的层数，和层序遍历的方式极其吻合。
// 在二叉树中，一层一层的来遍历二叉树，记录一下遍历的层数就是二叉树的深度
func maxDepth(root *TreeNode) int {
	depth := 0
	queue := make([]*TreeNode, 0)
	if root == nil {
		return depth
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		depth++
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return depth
}
