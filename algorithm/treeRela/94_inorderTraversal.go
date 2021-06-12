package treeRela

// 中序遍历递归实现
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	ret = append(ret, inorderTraversal(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversal(root.Right)...)
	return ret
}

// 中序遍历非递归实现，相当于自己模拟了一个栈
func inorderTraversalNotRecursion(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		// 将当前节点最左侧的所有节点入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 将整棵树左侧的节点从最左开始出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right
	}
	return ret
}
