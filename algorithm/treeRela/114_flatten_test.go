package treeRela

// 二叉树展开为链表
// 递归算法的关键要明确函数的定义，相信这个定义，而不要跳进递归细节！
// flatten 函数是怎么把左右子树拉平的？说不清楚，但是只要知道 flatten 的定义如此，相信这个定义，让 root 做它该做的事情，然后 flatten 函数就会按照定义工作。
// 注意：这里是后序遍历，必须要先将左右子树处理后再处理当前节点
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	// 保存原本的右子树
	beforeRight := root.Right
	// 将处理后的左子树作为新的右子树
	root.Left, root.Right = nil, root.Left
	// 将原本的右子树接到新右子树的末尾
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = beforeRight
}
