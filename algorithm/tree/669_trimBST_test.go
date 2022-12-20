package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//	            3
//	0                         4
//	        2
//	    1
//
// low = 1, high = 3
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 当前节点左子树都是不符合区间条件的，所以查找右子树，并将符合区间条件的返回
	if root.Val < low {
		// 以注释中的示例说明，这里当 root 是值为 0 的节点时，会 return 值为 2 的节点给上一层（值为 3 的节点），
		// 然后在之后会被 root.Left = trimBST(root.Left, low, high) 这部分代码将 3 的左子树替换掉，即删除了 0 的节点
		return trimBST(root.Right, low, high)
	}
	// 当前节点右子树都是不符合区间条件的，所以查找左子树，并将符合区间条件的返回
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// 当前节点在区间条件内，将左右节点替换为各自符合条件的子树
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// 迭代
func trimBSTNotRecursion(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到节点在范围的子树，由于范围外的节点要被 trim，所以直接用 root 来遍历
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low { // 当前节点的左子树不符合
			root = root.Right
		} else { // 当前节点的右子树不符合
			root = root.Left
		}
	}
	// 当前节点在区间内
	// 处理其左子树中不符合区间条件的节点
	cur := root
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}
	// 处理其右子树中不符合区间条件的节点，注意，这里需要将 cur 重置
	cur = root
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}
	return root
}

func TestTrimBST(t *testing.T) {
	tests := []struct {
		name        string
		levelOrders []interface{}
		low         int
		high        int
		want        []interface{}
	}{
		{
			name:        "first",
			levelOrders: []interface{}{1, 0, 2},
			low:         1,
			high:        2,
			want:        []interface{}{1, nil, 2},
		},
		{
			name:        "second",
			levelOrders: []interface{}{3, 0, 4, nil, 2, nil, nil, 1},
			low:         1,
			high:        3,
			want:        []interface{}{3, 2, nil, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root1 := constructTreeByArray(test.levelOrders)
			assert.Equal(t, test.levelOrders, root1.LevelPrint())
			assert.Equal(t, test.want, trimBST(root1, test.low, test.high).LevelPrint())

			root2 := constructTreeByArray(test.levelOrders)
			assert.Equal(t, test.want, trimBSTNotRecursion(root2, test.low, test.high).LevelPrint())
		})
	}
}
