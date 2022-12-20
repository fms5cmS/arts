package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		sum := 0
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			sum += cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 注意这里类型强转
		result = append(result, float64(sum)/float64(size))
	}
	return result
}

func TestAverageOfLevels(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  []float64
	}{
		{
			name:  "first",
			array: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want:  []float64{3.00000, 14.50000, 11.00000},
		},
		{
			name:  "second",
			array: []interface{}{3, 9, 20, 15, 7},
			want:  []float64{3.00000, 14.50000, 11.00000},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			assert.Equal(t, test.want, averageOfLevels(root))
		})
	}
}
