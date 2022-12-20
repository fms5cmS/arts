package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 中序遍历递归实现
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	ret = append(ret, inorderTraversal(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversal(root.Right)...)
	return ret
}

// 中序遍历非递归实现，相当于自己模拟了一个栈
//
//		   4
//		3    2
//	 1
//
// 最里层 for 循环先将最左侧节点入栈 stack = [4, 3, 1]
// 1 出栈，root = 1，并将值放入数组，array = 【1]，root=root.Right 为 nil
// root == nil，stack 不为空
// 3 出栈，root = 3，array = [1, 3]，root=root.Right 为 nil
// root == nil，stack 不为空
// 4 出栈，root = 4，array = [1, 3, 4]，root=root.Right 不为 nil
// root != nil，stack 为空
// 最里层 for 循环将 2 子树最左侧节点入栈，此时 root == nil
// 2 出栈，root = 2，array = [1, 3, 4, 2]，root=root.Right 为 nil
// root == nil，stack 为空，退出
func inorderTraversalByStack(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	// 指针的遍历来访问节点，栈用来处理节点上的元素
	for root != nil || len(stack) > 0 {
		// 将当前节点最左侧的所有节点入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 将整棵树左侧的节点从最左开始出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right
	}
	return ret
}

func TestInOrderTraversal(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  []int
	}{
		{
			name:  "first",
			array: []interface{}{1, nil, 2, 3},
			want:  []int{1, 3, 2},
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
			assert.Equal(t, test.want, inorderTraversal(root))
			assert.Equal(t, test.want, inorderTraversalByStack(root))
			assert.Equal(t, test.want, inorderTraversalNotRecursion(root))
		})
	}
}

// 遍历的非递归方法，统一代码
// 中序遍历的顺序是 左中右
// 入栈顺序是      右中左
func inorderTraversalNotRecursion(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil {
			// 右子节点入栈
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			// 注意这里会把非空的 cur 再次入栈，同时会入栈一个空指针来标记
			stack = append(stack, cur)
			stack = append(stack, nil)
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
