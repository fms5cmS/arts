package treeRela

import (
	"fmt"
	"testing"
)

func kthSmallest(root *TreeNode, k int) int {
	result := inorderForKthSmallest(root)
	// 这里 k <= len(result)，因为 k 是从 1 开始计数的
	if k > 0 && k <= len(result) {
		return result[k-1]
	}
	return 0
}

// 二叉搜索树的中序遍历是升序，所以查找第 k 小的元素，可以先通过中序遍历对 BST 排序
func inorderForKthSmallest(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	ret = append(ret, inorderForKthSmallest(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderForKthSmallest(root.Right)...)
	return ret
}

/*结果貌似不对？？*/

var (
	// 保存最后的结果
	target int
	// 记录遍历的节点个数
	count int
)

func kthSmallest2(root *TreeNode, k int) int {
	inorderPruningForKthSmallest(root, k)
	return target
}

func inorderPruningForKthSmallest(root *TreeNode, k int) {
	if root == nil {
		return
	}
	inorderPruningForKthSmallest(root.Left, k)
	// 剪枝操作
	count++
	if count == k {
		target = root.Val
		return
	}
	inorderPruningForKthSmallest(root.Right, k)
}

func TestKthSmallest(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1}},
			Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6},
	}
	//fmt.Println(kthSmallest(root,3))
	fmt.Println(kthSmallest2(root, 3))
}
