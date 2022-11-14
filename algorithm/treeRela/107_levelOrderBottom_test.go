package treeRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 在 102 的基础上反转结果集
func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0, size)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		result = append(result, tmp)
	}
	// 反转结果集
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}
	return result
}

func TestLevelOrderBottom(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  [][]int
	}{
		{
			name:  "first",
			array: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want:  [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			name:  "second",
			array: []interface{}{1},
			want:  [][]int{{1}},
		},
		{
			name:  "third",
			array: []interface{}{},
			want:  [][]int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			assert.Equal(t, test.want, levelOrderBottom(root))
		})
	}
}
