package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findBottomLeftValue(root *TreeNode) int {
	queue, result := make([]*TreeNode, 0), 0
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 记录最后一行第一个元素
			if i == 0 {
				result = cur.Val
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return result
}

func TestFindBottomLeftValue(t *testing.T) {
	tests := []struct {
		name      string
		leveOrder []interface{}
		want      int
	}{
		{
			name:      "first",
			leveOrder: []interface{}{2, 1, 3},
			want:      1,
		},
		{
			name:      "second",
			leveOrder: []interface{}{1, 2, 3, 4, nil, 5, 6, nil, nil, 7},
			want:      7,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.leveOrder)
			assert.Equal(t, test.leveOrder, root.LevelPrint())

			assert.Equal(t, test.want, findBottomLeftValue(root))
		})
	}
}
