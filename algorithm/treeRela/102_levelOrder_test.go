package treeRela

import (
	"fmt"
	"testing"
)

// 层序遍历：队列实现
// 广度优先遍历
// 除了下面这样的写法，也可以像 base.go 中那样每层使用一个变量来保存节点队列
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0)
		// 这里要使用固定的 size，因为 queue 的长度是变化的，而 size 代表了这一层的节点数
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		result = append(result, tmp)
	}
	return result
}

func TestLevelOrder2(t *testing.T) {
	a := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	ret := levelOrder(a)
	for _, ints := range ret {
		fmt.Println(ints)
	}
}
