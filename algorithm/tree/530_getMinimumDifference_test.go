package tree

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 注意题目说的是二叉搜索树，而二叉搜索树是有序的！
// 所以可以当作是在一个有序数组中求解两个数的最小差值
// 最直观的想法，就是把二叉搜索树转换成有序数组，然后遍历一遍数组，就统计出来最小差值了。
func getMinimumDifference(root *TreeNode) int {
	// 中序遍历得到数组
	array := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		array = append(array, root.Val)
		root = root.Right
	}
	// 遍历数组找到最小差值
	min := math.MaxInt64
	for i := 0; i < len(array)-1; i++ {
		if array[i+1]-array[i] < min {
			min = array[i+1] - array[i]
		}
	}
	return min
}

func getMinimumDifference2(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	min := math.MaxInt64
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil {
				if cur.Val-pre.Val < min {
					min = cur.Val - pre.Val
				}
			}
			pre = cur
			cur = cur.Right
		}
	}
	return min
}

func TestGetMinimumDifference(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  int
	}{
		{
			name:  "first",
			array: []interface{}{4, 2, 6, 1, 3},
			want:  1,
		},
		{
			name:  "second",
			array: []interface{}{1, 0, 48, nil, nil, 12, 49},
			want:  1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			assert.Equal(t, test.array, root.LevelPrint())

			assert.Equal(t, test.want, getMinimumDifference(root))
			assert.Equal(t, test.want, getMinimumDifference2(root))
		})
	}
}
