package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 计算给定二叉树的所有**左叶子**之和。
// 判断当前节点是不是左叶子是无法判断的，必须要通过节点的父节点来判断其左孩子是不是左叶子!
// 递归
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	midValue := 0
	// 找到左叶子节点
	if root.Left != nil && (root.Left.Left == nil && root.Left.Right == nil) {
		midValue = root.Left.Val
	}
	leftValue := sumOfLeftLeaves(root.Left)   // 左子树左叶子之和
	rightValue := sumOfLeftLeaves(root.Right) // 右子树左叶子之和
	return leftValue + rightValue + midValue
}

// 迭代法，前中后序遍历都可以，这里使用前序遍历
func sumOfLeftLeaves2(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	if root == nil {
		return 0
	}
	stack = append(stack, root)
	result := 0
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Left != nil && cur.Left.Left == nil && cur.Left.Right == nil {
			result += cur.Left.Val
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return result
}

func TestSumOfLeftLeavesr(t *testing.T) {
	tests := []struct {
		name      string
		leveOrder []interface{}
		want      int
	}{
		{
			name:      "first",
			leveOrder: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want:      24,
		},
		{
			name:      "second",
			leveOrder: []interface{}{1},
			want:      0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.leveOrder)
			assert.Equal(t, test.leveOrder, root.LevelPrint())

			assert.Equal(t, test.want, sumOfLeftLeaves(root))
			assert.Equal(t, test.want, sumOfLeftLeaves2(root))
		})
	}
}
