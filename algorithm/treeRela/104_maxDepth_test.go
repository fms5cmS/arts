package treeRela

import (
	"arts/algorithm/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 迭代法的话，使用层序遍历是最为合适的，因为最大的深度就是二叉树的层数，和层序遍历的方式极其吻合。
// 在二叉树中，一层一层的来遍历二叉树，记录一下遍历的层数就是二叉树的深度
func maxDepth(root *TreeNode) int {
	depth := 0
	queue := make([]*TreeNode, 0)
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

// 递归求树的最大深度
// 前序遍历求的是树深度，后序遍历求的是树的高度。这道题两种遍历方式都可以，这里采用的是后序遍历
// 根节点的高度就是树的最大深度
func maxDepthByRecursion1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepthByRecursion1(root.Left)
	rightDepth := maxDepthByRecursion1(root.Right)
	return 1 + utils.Max(leftDepth, rightDepth)
}

// 前序遍历求树的最大深度
func maxDepthByRecursion2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getDepthForMax(root, 1, 0)
}

// 查询指定节点的深度
func getDepthForMax(root *TreeNode, depth int, result int) int {
	if depth > result {
		result = depth // 中
	}
	if root.Left == nil && root.Right == nil {
		return result
	}
	if root.Left != nil { // 左
		result = getDepthForMax(root.Left, depth+1, result)
	}
	if root.Right != nil { // 右
		result = getDepthForMax(root.Right, depth+1, result)
	}
	return result
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  int
	}{
		{
			name:  "first",
			array: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want:  3,
		},
		{
			name:  "second",
			array: []interface{}{1, nil, 2},
			want:  2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			assert.Equal(t, test.want, maxDepth(root))
			assert.Equal(t, test.want, maxDepthByRecursion1(root))
			assert.Equal(t, test.want, maxDepthByRecursion2(root))
		})
	}
}
