package treeRela

import "math"

func isBalanced(root *TreeNode) bool {
	if getHeight(root) == -1 {
		return false
	}
	return true
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left) // 左
	if leftHeight == -1 {
		return -1 // 说明左子树已经不是二叉平衡树
	}
	rightHeight := getHeight(root.Right) // 右
	if rightHeight == -1 {
		return -1 // 说明右子树已经不是二叉平衡树
	}
	// 计算左右子树的高度差
	if math.Abs(float64(leftHeight-rightHeight)) > 1 { // 中
		return -1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// 以当前节点为根节点的最大高度
	return 1 + max(leftHeight, rightHeight)
}
