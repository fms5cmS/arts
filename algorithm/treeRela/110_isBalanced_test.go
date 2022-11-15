package treeRela

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 平衡树需要左右子树的高度相等，所以需要计算树的高度
// 由于要比较左右子树的高度，所以必然是后序遍历
func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left) // 左
	if leftHeight == -1 {
		return -1 // 说明左子树已经不是二叉平衡树
	}
	rightHeight := getHeight(root.Right) // 右
	if rightHeight == -1 {
		return -1 // 说明右子树已经不是二叉平衡树
	}
	// 计算左右子树的高度差
	// 平衡树要求两个子树的高度差的绝对值不超过 1
	if math.Abs(float64(leftHeight-rightHeight)) > 1 { // 中
		return -1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// 以当前节点为根节点的最大高度
	return 1 + max(leftHeight, rightHeight)
}

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  bool
	}{
		{
			name:  "first",
			array: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want:  true,
		},
		{
			name:  "second",
			array: []interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4},
			want:  false,
		},
		{
			name:  "third",
			array: []interface{}{},
			want:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			assert.Equal(t, test.want, isBalanced(root))
		})
	}
}
