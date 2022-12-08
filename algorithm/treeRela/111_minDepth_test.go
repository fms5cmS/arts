package treeRela

import (
	"arts/algorithm/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
func minDepth(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	depth := 0
	if root == nil {
		return depth
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		depth++
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 只有当左右孩子都为空的时候，才说明遍历到最低点了。如果其中一个孩子为空则不是最低点
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return depth
}

// 后序遍历求最小深度
func minDepthByRecursion(root *TreeNode) int {
	return getDepthForMin(root)
}

func getDepthForMin(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getDepthForMin(root.Left)
	rightDepth := getDepthForMin(root.Right)
	// 由于最小深度是从根节点到最近**叶子节点**的最短路径上的节点数量。
	// 所以如果只有一个子节点为空时，此时并不是最低点
	if root.Left == nil && root.Right != nil {
		return 1 + rightDepth
	}
	if root.Left != nil && root.Right == nil {
		return 1 + leftDepth
	}
	// 走到这里说明到了叶子节点
	return 1 + utils.Min(leftDepth, rightDepth)
}

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  int
	}{
		{
			name:  "first",
			array: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want:  2,
		},
		{
			name:  "second",
			array: []interface{}{2, nil, 3, nil, 4, nil, 5, nil, 6},
			want:  5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			assert.Equal(t, test.want, minDepth(root))
			assert.Equal(t, test.want, minDepthByRecursion(root))
		})
	}
}
