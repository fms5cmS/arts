package treeRela

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 递归函数，累加的方式代码比较复杂，所以使用递减
	// 计数器 count 初始值为 targetSum，每次减去节点上的数值
	var traversal func(cur *TreeNode, count int) bool
	traversal = func(cur *TreeNode, count int) bool {
		//  遇到叶子节点，且此时计数为 0
		if cur.Left == nil && cur.Right == nil && count == 0 {
			return true
		}
		if cur.Left == nil && cur.Right == nil {
			return false
		}
		// 处理左子树
		if cur.Left != nil {
			count -= cur.Left.Val
			if traversal(cur.Left, count) {
				return true
			}
			count += cur.Left.Val // 回溯
		}
		// 处理右子树
		if cur.Right != nil {
			count -= cur.Right.Val
			if traversal(cur.Right, count) {
				return true
			}
			count += cur.Right.Val // 回溯
		}
		return false
	}
	return traversal(root, targetSum-root.Val)
}

func hasPathSumNoRecursion(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	type NodeSum struct {
		Node *TreeNode // 节点指针
		Sum  int       // 路径数值之和
	}
	stack := make([]*NodeSum, 0)
	stack = append(stack, &NodeSum{Node: root, Sum: root.Val})
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Node.Left == nil && cur.Node.Right == nil && cur.Sum == targetSum {
			return true
		}
		if cur.Node.Right != nil {
			stack = append(stack, &NodeSum{Node: cur.Node.Right, Sum: cur.Sum + cur.Node.Right.Val})
		}
		if cur.Node.Left != nil {
			stack = append(stack, &NodeSum{Node: cur.Node.Left, Sum: cur.Sum + cur.Node.Left.Val})
		}
	}
	return false
}
