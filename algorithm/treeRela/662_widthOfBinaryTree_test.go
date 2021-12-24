package treeRela

import (
	"fmt"
	"testing"
)

type item struct {
	idx       int // 如果以数组存储整棵树时，该节点对应的索引（从 0 开始）
	*TreeNode     // 二叉树的节点
}

// 注意：每一层的最大宽度是指左右两侧均为非 nil 的节点之间的宽度！
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 默认宽度为 1，即根节点所在的第一层
	maxWidth, queue := 1, make([]item, 0)
	// 根节点的索引为 0
	queue = append(queue, item{0, root})
	// 层序遍历
	for len(queue) > 0 {
		size := len(queue)
		// 计算每一层的宽度 width，并比较其与 maxWidth 的大小
		if width := queue[size-1].idx - queue[0].idx + 1; width > maxWidth {
			maxWidth = width
		}
		// 向队列中添加下一层的每个节点及其对应的索引
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, item{cur.idx*2 + 1, cur.Left})
			}
			if cur.Right != nil {
				queue = append(queue, item{cur.idx*2 + 2, cur.Right})
			}
		}
	}
	return maxWidth
}

func TestMaxWidth(t *testing.T) {
	a := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 9, Right: &TreeNode{Val: 7}}},
	}
	// a := &TreeNode{
	// 	Val: 1,
	// 	Right: &TreeNode{Val: 2},
	// }
	// a := &TreeNode{
	// 	Val:  1,
	// 	Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3}},
	// }
	fmt.Println(widthOfBinaryTree(a))
}
