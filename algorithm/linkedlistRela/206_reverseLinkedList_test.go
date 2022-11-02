package linkedlistRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 迭代解法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = nil
	for head != nil {
		tmp := head.Next
		head.Next = pre // 反转
		pre, head = head, tmp
	}
	return pre
}

// 递归解法
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return reverse(nil, head)
}

func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	tmp := cur.Next
	cur.Next = pre
	return reverse(cur, tmp)
}

func TestReverse(t *testing.T) {
	tests := []struct {
		array  []int
		output []int
	}{
		{
			array:  []int{1, 2, 3, 4, 5},
			output: []int{5, 4, 3, 2, 1},
		},
		{
			array:  []int{1, 2},
			output: []int{2, 1},
		},
		{
			array:  []int{1},
			output: []int{1},
		},
		{
			array:  []int{},
			output: []int{},
		},
	}
	for _, test := range tests {
		before := generateListViaArray(test.array)
		after := reverseList(before)
		assert.Equal(t, test.output, generateArrayViaList(after))
	}
}
