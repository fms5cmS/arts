package treeRela

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

// 遍历的非递归方法，统一代码
func PostorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil {
			// 注意这里会把非空的 cur 再次入栈，同时会入栈一个空指针来标记
			stack = append(stack, cur)
			stack = append(stack, nil)
			// 右子节点入栈
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
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

// 遍历的非递归方法，统一代码
// 后序遍历的顺序是 左右中
// 入栈顺序是      中右左
func postorderTraversalNotRecursion(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}
	// 上面向 result 添加元素的顺序为 中、右、左，要得到后序遍历的 左、右、中，需要反转切片元素！！
	reverseInts(result)
	return result
}

// 注意，这里修改了底层数组，所以不需要返回切片了
// 注意与 151 中的区别！
func reverseInts(src []int) {
	left, right := 0, len(src)-1
	for left < right {
		src[left], src[right] = src[right], src[left]
		left++
		right--
	}
}
