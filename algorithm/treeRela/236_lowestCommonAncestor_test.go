package treeRela

// 如何找到 p、q 最近的公共祖先呢？ ——> 从下往上查找即可 ——> 即回溯，而二叉树的回溯就是后序遍历
// 而后序遍历可以使用递归来实现，每次递归找到该子树中的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
