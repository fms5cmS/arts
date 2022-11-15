package treeRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 层序遍历属于广度优先遍历：队列实现
// 除了下面这样的写法，也可以像 base.go 中那样每层使用一个变量来保存节点队列
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // 需要记录当前层的节点数量
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

func TestLevelOrder2(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  [][]int
	}{
		{
			name:  "first",
			array: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want:  [][]int{{3}, {9, 20}, {15, 7}},
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
			assert.Equal(t, test.want, levelOrder(root))
		})
	}
}
