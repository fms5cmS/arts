package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index, max := maxOfNums(nums)
	root := new(TreeNode)
	root.Val = max
	// 左闭右开区间，所以这里是 index
	root.Left = constructMaximumBinaryTree(nums[:index])
	if index < len(nums) {
		root.Right = constructMaximumBinaryTree(nums[index+1:])
	}
	return root
}

func maxOfNums(nums []int) (index, max int) {
	for i, num := range nums {
		if num > max {
			index, max = i, num
		}
	}
	return
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []interface{}
	}{
		{
			name: "first",
			nums: []int{3, 2, 1, 6, 0, 5},
			want: []interface{}{6, 3, 5, nil, 2, 0, nil, nil, 1},
		},
		{
			name: "second",
			nums: []int{3, 2, 1},
			want: []interface{}{3, nil, 2, nil, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructMaximumBinaryTree(test.nums)
			assert.Equal(t, test.want, root.LevelPrint())
		})
	}
}
