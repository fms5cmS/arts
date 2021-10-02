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

// 遍历的非递归方法，统一代码
// 中序遍历的顺序是 左中右
// 入栈顺序是      右中左
func InorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil {
			// 右子节点入栈
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			// 注意这里会把非空的 cur 再次入栈，同时会入栈一个空指针来标记
			stack = append(stack, cur)
			stack = append(stack, nil)
			// 左子节点入栈
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		} else {
			// cur 为 nil，代表当前栈顶元素为中间节点
			// 取出栈顶元素，并添加到 result
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, cur.Val)
		}
	}
	return result
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
