package treeRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

func searchBSTNotRecursion(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

func TestSearchBST(t *testing.T) {
	tests := []struct {
		name        string
		levelOrders []interface{}
		val         int
		want        []interface{}
	}{
		{
			name:        "first",
			levelOrders: []interface{}{4, 2, 7, 1, 3},
			val:         2,
			want:        []interface{}{2, 1, 3},
		},
		{
			name:        "second",
			levelOrders: []interface{}{4, 2, 7, 1, 3},
			val:         5,
			want:        []interface{}{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.levelOrders)
			get1 := searchBST(root, test.val)
			get2 := searchBSTNotRecursion(root, test.val)

			assert.Equal(t, test.want, get1.LevelPrint())
			assert.Equal(t, test.want, get2.LevelPrint())
		})
	}
}
