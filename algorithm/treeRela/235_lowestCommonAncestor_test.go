package treeRela

// 二叉搜索树的最近公共祖先
// 给定一个二叉搜索树, 找到**该树中两个指定节点**的最近公共祖先。
// 二叉搜索树是有序的，如何利用呢？
// 只要从上到下遍历的时候，cur 节点是数值在 p、q 区间中则说明 cur 是最近公共祖先了
func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	// 该终止条件可以去掉，题目中说了p、q 为不同节点且均存在于给定的二叉搜索树中。也就是说一定会找到公共祖先的，所以并不存在遇到空的情况。
	if root == nil {
		return root
	}
	// 目标区间在左子树，注意：p、q 的值谁大并不知道，所以这里两个值都要判断
	if root.Val > p.Val && root.Val > q.Val {
		if left := lowestCommonAncestor235(root.Left, p, q); left != nil {
			// 注意与 236 的区别，236 查找完一条边后不能直接 return
			// 236 中是普通的二叉树，需要从下往上回溯，所以需要遍历整棵树并对 left、right 再进行逻辑判断
			// 235 这里利用了 BST 的特性，使用的是从上往下的遍历，只要树的单边满足了区间的要求，就说明得到目标节点了
			return left // 直接 return 了！
		}
	}
	// 目标区间在右子树
	if root.Val < p.Val && root.Val < q.Val {
		if right := lowestCommonAncestor235(root.Right, p, q); right != nil {
			return right
		}
	}
	// 目标区间在 p、q 之间，直接返回当前节点
	return root
}

// 迭代
func lowestCommonAncestor235NotRecursion(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}