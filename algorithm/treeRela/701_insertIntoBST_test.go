package treeRela

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	// 注意：需要将插入后的根节点保存为 root 对应的子树！
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

// 迭代
func insertIntoBSTNotRecursion(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	// parent 用于记录要插入位置的父节点
	cur, parent := root, root
	// 遍历找到插入的位置
	for cur != nil {
		parent = cur
		if val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	// 插入
	if val < parent.Val {
		parent.Left = node
	} else {
		parent.Right = node
	}
	return root
}
