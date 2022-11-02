package linkedlistRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 迭代解法
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	pre := dummy
	for head != nil && head.Next != nil {
		first, second := head, head.Next
		pre.Next = second
		// 交换相邻两个节点
		first.Next, second.Next = second.Next, first
		pre, head = first, first.Next
	}
	return dummy.Next
}

// 递归解法
func swapPairsRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	first.Next = swapPairsRecursion(second.Next)
	second.Next = first
	return second
}

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		array  []int
		output []int
	}{
		{
			array:  []int{1, 2, 3, 4},
			output: []int{2, 1, 4, 3},
		},
		{
			array:  []int{},
			output: []int{},
		},
		{
			array:  []int{1},
			output: []int{1},
		},
		{
			array:  []int{1, 2, 3},
			output: []int{2, 1, 3},
		},
	}
	for _, test := range tests {
		before := generateListViaArray(test.array)
		after := swapPairs(before)
		assert.Equal(t, test.output, generateArrayViaList(after), generateArrayViaList(swapPairsRecursion(before)))
	}
}
