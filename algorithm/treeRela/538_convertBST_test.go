package treeRela

// BST 的中序遍历是升序的，那么反过来遍历得到的就是降序的
// 先遍历右子树后遍历根节点再遍历左子树
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 注意这里需要有一个外部变量记录和
	sum := 0
	var walk func(root *TreeNode)
	walk = func(root *TreeNode) {
		if root == nil {
			return
		}
		walk(root.Right)
		sum += root.Val
		root.Val = sum
		walk(root.Left)
	}
	walk(root)
	return root
}
