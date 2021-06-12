package treeRela

import "testing"

// 前序和后序遍历的方式可以满足，中序则不行！
// 中序遍历时，先翻转了 root 的左子树 left，然后会把 left 和 right 翻转，再翻转 root 的右子树，而此时的右子树是原本的左子树了，再次翻转会被还原
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 前序遍历的方式
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	// 后序遍历的方式
	// invertTree(root.Left)
	// invertTree(root.Right)
	// root.Left, root.Right = root.Right, root.Left
	// 中序遍历的方式无法实现翻转
	// invertTree(root.Left)
	// root.Left, root.Right = root.Right, root.Left
	// invertTree(root.Right)
	return root
}

func TestInvertTree(t *testing.T) {
	a := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}},
	}
	aTree := LevelOrder(a)
	for _, arr := range aTree {
		t.Log(arr)
	}
	t.Log("翻转二叉树后")
	tree := invertTree(a)
	treeArr := LevelOrder(tree)
	for _, arr := range treeArr {
		t.Log(arr)
	}
}
