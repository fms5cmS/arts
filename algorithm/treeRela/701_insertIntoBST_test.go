package treeRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 这些都是不改变树的结构，将新的值往叶子节点插入的！！！

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	// 注意：需要将插入后的根节点保存为 root 对应的子树！
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

// 迭代
func insertIntoBSTNotRecursion(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	// parent 用于记录要插入位置的父节点
	cur, parent := root, root
	// 遍历找到插入的位置
	for cur != nil {
		parent = cur
		if val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	// 插入
	if val < parent.Val {
		parent.Left = node
	} else {
		parent.Right = node
	}
	return root
}

func TestInsertIntoBST(t *testing.T) {
	tests := []struct {
		name        string
		levelOrders []interface{}
		val         int
		want        []interface{}
	}{
		{
			name:        "first",
			levelOrders: []interface{}{4, 2, 7, 1, 3},
			val:         5,
			want:        []interface{}{4, 2, 7, 1, 3, 5},
		},
		{
			name:        "second",
			levelOrders: []interface{}{40, 20, 60, 10, 30, 50, 70},
			val:         25,
			want:        []interface{}{40, 20, 60, 10, 30, 50, 70, nil, nil, 25},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root1 := constructTreeByArray(test.levelOrders)
			assert.Equal(t, test.levelOrders, root1.LevelPrint())

			assert.Equal(t, test.want, insertIntoBST(root1, test.val).LevelPrint())

			root2 := constructTreeByArray(test.levelOrders)
			assert.Equal(t, test.want, insertIntoBSTNotRecursion(root2, test.val).LevelPrint())
		})
	}
}
