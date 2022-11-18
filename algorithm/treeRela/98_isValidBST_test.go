package treeRela

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 根据 BST 的特性（左子树的值小于根节点，根节点的值小于右子树）来判断
// 注意：这里要小心存在特殊情况，如下面的树并不是 BST，因为 6 小于 10，但却位于 10 的右子树中
//
//	    10
//	5       15
//	       6   20
//
// 所以需要得到每棵子树的最小最大值来比较判断
func isValidBST(root *TreeNode) bool {
	// 在遍历过程中就完成数组有序性的比较
	// min 是右子树的最小值，max 是左子树的最大值
	var isValid func(root *TreeNode, min, max int) bool
	isValid = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}
		return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
	}
	return isValid(root, math.MinInt64, math.MaxInt64)
}

// BST 中序遍历后得到升序的节点值
// 先中序遍历，然后判断是否为有序的即可
func isValidBST2(root *TreeNode) bool {
	arr := getInorderArr(root)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}

func getInorderArr(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	tmp := make([]int, 0)
	tmp = append(tmp, getInorderArr(root.Left)...)
	tmp = append(tmp, root.Val)
	tmp = append(tmp, getInorderArr(root.Right)...)
	return tmp
}

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  bool
	}{
		{
			name:  "first",
			array: []interface{}{2, 1, 3},
			want:  true,
		},
		{
			name:  "second",
			array: []interface{}{5, 1, 4, nil, nil, 3, 6},
			want:  false,
		},
		{
			name:  "third",
			array: []interface{}{2, 2, 2},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			assert.Equal(t, test.array, root.LevelPrint())

			assert.Equal(t, test.want, isValidBST(root))
			assert.Equal(t, test.want, isValidBST2(root))
		})
	}
}
