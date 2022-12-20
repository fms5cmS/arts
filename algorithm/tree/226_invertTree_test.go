package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 前序和后序遍历的方式可以满足，中序则不行！
// 中序遍历时，先翻转了 root 的左子树 left，然后会把 left 和 right 翻转，再翻转 root 的右子树，而此时的右子树是原本的左子树了，再次翻转会被还原
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 前序遍历的方式
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	// 后序遍历的方式
	// invertTree(root.Left)
	// invertTree(root.Right)
	// root.Left, root.Right = root.Right, root.Left
	// 中序遍历的方式无法实现翻转
	// invertTree(root.Left)
	// root.Left, root.Right = root.Right, root.Left
	// invertTree(root.Right)
	return root
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  []int
	}{
		{
			name:  "first",
			array: []interface{}{4, 2, 7, 1, 3, 6, 9},
			want:  []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:  "second",
			array: []interface{}{2, 1, 3},
			want:  []int{2, 3, 1},
		},
		{
			name:  "third",
			array: []interface{}{},
			want:  []int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			after := invertTree(root)
			afterLevel := LevelOrder(after)
			get := make([]int, 0)
			for _, ints := range afterLevel {
				get = append(get, ints...)
			}
			assert.Equal(t, test.want, get)
		})
	}
}
