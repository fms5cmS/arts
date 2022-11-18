package treeRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// BST 的题目，可以看作是对有序数组的操作，也就是对有序数组从后往前累加
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

func convertBSTNotRecursion(root *TreeNode) *TreeNode {
	// 记录前一个节点的值
	pre := 0
	traversal := func(root *TreeNode) {
		stack := make([]*TreeNode, 0)
		cur := root
		for cur != nil || len(stack) > 0 {
			if cur != nil {
				stack = append(stack, cur)
				cur = cur.Right
			} else {
				cur = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				cur.Val += pre
				pre = cur.Val
				cur = cur.Left
			}
		}
	}
	traversal(root)
	return root
}

func TestConvertBST(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  []interface{}
	}{
		{
			name:  "first",
			array: []interface{}{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8},
			want:  []interface{}{30, 36, 21, 36, 35, 26, 15, nil, nil, nil, 33, nil, nil, nil, 8},
		},
		{
			name:  "second",
			array: []interface{}{0, nil, 1},
			want:  []interface{}{1, nil, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root1 := constructTreeByArray(test.array)
			assert.Equal(t, test.array, root1.LevelPrint())

			assert.Equal(t, test.want, convertBST(root1).LevelPrint())

			root2 := constructTreeByArray(test.array)
			assert.Equal(t, test.want, convertBST(root2).LevelPrint())
		})
	}
}
