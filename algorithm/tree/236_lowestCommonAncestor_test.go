package tree

// 二叉树的最近公共祖先
// 如何找到 p、q 最近的公共祖先呢？ ——> 从下往上查找即可 ——> 即回溯，而二叉树的回溯就是后序遍历
// 而后序遍历可以使用递归来实现，每次递归找到该子树中的最近公共祖先
// 如果判断某一节点是 p、q 的公共祖先呢？其左右子树各自出现了 p、q 节点，则说明该节点是公共祖先
func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	// 分别找左右子树中 p、q 的公共祖先
	// 即使已经找到结果了，依然要把其他节点遍历完，因为要使用递归函数的返回值（left、right）做逻辑判断
	left := lowestCommonAncestor236(root.Left, p, q)
	right := lowestCommonAncestor236(root.Right, p, q)
	// left、right 都不为空，则说明 root 是最近的公共祖先
	// left 和 right 分别为 p、q 会返回自身
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
