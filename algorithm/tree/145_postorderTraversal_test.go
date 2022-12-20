package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

// 遍历的非递归方法，统一代码
// 后序遍历的顺序是 左右中
// 先得到 中右左，然后反转
func postorderTraversalByStack(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}
	// 上面向 result 添加元素的顺序为 中、右、左，要得到后序遍历的 左、右、中，需要反转切片元素！！
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}
	return result
}

func TestPostOrder(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  []int
	}{
		{
			name:  "first",
			array: []interface{}{1, nil, 2, 3},
			want:  []int{3, 2, 1},
		},
		{
			name:  "second",
			array: []interface{}{},
			want:  []int{},
		},
		{
			name:  "third",
			array: []interface{}{1},
			want:  []int{1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			assert.Equal(t, test.want, postorderTraversal(root))
			assert.Equal(t, test.want, postorderTraversalByStack(root))
			assert.Equal(t, test.want, postorderTraversalNotRecursion(root))
		})
	}
}

// 遍历的非递归方法，统一代码
func postorderTraversalNotRecursion(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil {
			// 注意这里会把非空的 cur 再次入栈，同时会入栈一个空指针来标记
			stack = append(stack, cur)
			stack = append(stack, nil)
			// 右子节点入栈
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			// 左子节点入栈
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		} else {
			// cur 为 nil，代表当前栈顶元素为中间节点
			// 取出栈顶元素，并添加到 result
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, cur.Val)
		}
	}
	return result
}
