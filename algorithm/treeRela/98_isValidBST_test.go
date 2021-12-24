package treeRela

import (
	"math"
)

// 根据 BST 的特性（左子树的值小于根节点，根节点的值小于右子树）来判断
// 注意：这里要小心存在特殊情况，如下面的树并不是 BST，因为 6 小于 10，但却位于 10 的右子树中
//       10
//   5       15
//          6   20
// 所以需要得到每棵子树的最小最大值来比较判断
func isValidBST(root *TreeNode) bool {
	var isValid func(root *TreeNode, min, max int) bool
	isValid = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}
		return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
	}
	return isValid(root, math.MinInt64, math.MaxInt64)
}

// BST 中序遍历后得到升序的节点值
// 先中序遍历，然后判断是否为有序的即可
func isValidBST2(root *TreeNode) bool {
	arr := inorderValid(root)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}

func inorderValid(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	tmp := make([]int, 0)
	tmp = append(tmp, inorderValid(root.Left)...)
	tmp = append(tmp, root.Val)
	tmp = append(tmp, inorderValid(root.Right)...)
	return tmp
}
