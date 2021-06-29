package treeRela

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := root.Left, root.Right
	// 注意这里要初始化为 0，为了在满二叉树时方便计算
	leftHeight, rightHeight := 0, 0
	// 分别计算 root 的左右子树的的高度
	// 注意：这里只需计算了最左和最右侧的叶子节点高度来判断这棵完全二叉树是否为满二叉树
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

func countNodes_Easy(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 后序遍历
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
