package treeRela

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

// 遍历的非递归方法，统一代码
// 前序遍历的顺序是 中左右
// 入栈顺序是      右左中
func PreorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0{
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil {
			// 右子节点入栈
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			// 左子节点入栈
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			// 注意这里会把非空的 cur 再次入栈，同时会入栈一个空指针来标记
			stack = append(stack, cur)
			stack = append(stack, nil)
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

func preorderTraversalNotRecursion(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		// 从数组尾部开始出栈！
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return result
}
