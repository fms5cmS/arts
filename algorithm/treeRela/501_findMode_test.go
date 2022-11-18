package treeRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 中序遍历是为了利用 BST 中序遍历结果有序的特征
// 有序数组中，相同元素一定是相邻的
func findMode(root *TreeNode) []int {
	// count 记录当前元素出现的频率，maxCount 记录整棵树中元素出现的最大频率
	count, maxCount := 0, 0
	result := make([]int, 0)
	var pre *TreeNode
	// 中序遍历
	var searchBSTForMode func(node *TreeNode)
	searchBSTForMode = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		searchBSTForMode(cur.Left)
		if pre == nil {
			count = 1 // 当前节点的值 count 更新为 1
		} else if pre.Val == cur.Val {
			count++
		} else {
			count = 1 // 当前节点的值 count 更新为 1，如果不是遍历顺序不是有序的，就不能直接更新了，而要采用 map 统计每个元素的出现频率！
		}
		pre = cur
		if count == maxCount {
			result = append(result, cur.Val)
		}
		// count 大于 maxCount 时，需要重置 maxCount 值以及 result
		if count > maxCount {
			maxCount = count
			result = []int{cur.Val}
		}
		searchBSTForMode(cur.Right)
	}
	searchBSTForMode(root)
	return result
}

// 迭代
func findModeNotRecursion(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	var (
		cur           = root
		pre *TreeNode = nil
	)
	count, maxCount := 0, 0
	result := make([]int, 0)
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left // 左
		} else {
			cur = stack[len(stack)-1] // 中
			stack = stack[:len(stack)-1]
			if pre == nil {
				count = 1
			} else if pre.Val == cur.Val {
				count++
			} else {
				count = 1
			}
			if count == maxCount {
				result = append(result, cur.Val)
			}
			// count 大于 maxCount 时，需要重置 maxCount 值以及 result
			if count > maxCount {
				maxCount = count
				result = []int{cur.Val}
			}
			pre = cur
			cur = cur.Right // 右
		}
	}
	return result
}

func findModeSelf(root *TreeNode) []int {
	// 中序遍历得到数组
	array := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		array = append(array, cur.Val)
		cur = cur.Right
	}
	// 找到数组中重复出现的元素
	maxCount := 0
	result := make([]int, 0)
	slow, fast := 0, 0
	for fast < len(array) {
		for fast < len(array) && array[fast] == array[slow] {
			fast++
		}
		count := fast - slow
		if count > maxCount {
			maxCount = count
			result = []int{array[slow]} // 清空并放入新的元素
		} else if count == maxCount {
			result = append(result, array[slow]) // 追加元素
		}
		slow = fast
	}
	return result
}

func TestFindMode(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  []int
	}{
		{
			name:  "first",
			array: []interface{}{1, nil, 2, 2},
			want:  []int{2},
		},
		{
			name:  "second",
			array: []interface{}{0},
			want:  []int{0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			assert.Equal(t, test.array, root.LevelPrint())

			assert.Equal(t, test.want, findMode(root))
			assert.Equal(t, test.want, findModeNotRecursion(root))
			assert.Equal(t, test.want, findModeSelf(root))
		})
	}
}
