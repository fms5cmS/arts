package treeRela

// 后序遍历
func countNodesByRecursion(root *TreeNode) int {
	return getNodeNum(root)
}

func getNodeNum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftNum := getNodeNum(root.Left)   // 左
	rightNum := getNodeNum(root.Right) // 右
	treeNum := leftNum + rightNum + 1  // 中
	return treeNum
}

// 利用了完全二叉树的性质
// 完全二叉树只有两种情况：
// 情况一，就是满二叉树。可以用 2^深度-1 计算，注意，根节点深度为 1
// 情况二，最后一层叶子节点没有满。分别递归左孩子，和右孩子，递归到某一深度一定会有左孩子或者右孩子为满二叉树，然后依然可以按照情况1来计算。
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := root.Left, root.Right
	// 注意这里要初始化为 0，为了在满二叉树时方便计算
	leftHeight, rightHeight := 0, 0
	// 分别计算 root 的左右子树的的高度
	// 注意：这里只需计算了最左和最右侧的叶子节点高度来判断这棵完全二叉树是否为满二叉树！！
	for left != nil {
		left = left.Left
		leftHeight++
	}
	for right != nil {
		right = right.Right
		rightHeight++
	}
	// 左右子树高度相等，则 root 这棵树就是满二叉树了，可以使用公式计算该树的节点数
	if leftHeight == rightHeight {
		// Go 中 >>、<<、&、&^ 这四个位运算符的优先级是高于 +、- 的，这与其他语言不同
		return 2<<leftHeight - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// 非递归的方式可以使用层序遍历
