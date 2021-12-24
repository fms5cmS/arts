package treeRela

//                    3
//        0                         4
//                2
//            1
//
// low = 1, high = 3
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 当前节点左子树都是不符合区间条件的，所以查找右子树，并将符合区间条件的返回
	if root.Val < low {
		// 以注释中的示例说明，这里当 root 是值为 0 的节点时，会 return 值为 2 的节点给上一层（值为 3 的节点），
		// 然后在之后会被 root.Left = trimBST(root.Left, low, high) 这部分代码将 3 的左子树替换掉，即删除了 0 的节点
		return trimBST(root.Right, low, high)
	}
	// 当前节点右子树都是不符合区间条件的，所以查找左子树，并将符合区间条件的返回
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// 当前节点在区间条件内，将左右节点替换为各自符合条件的子树
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}


//func trimBSTNotRecursion(root *TreeNode, low int, high int) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	for root != nil
//}