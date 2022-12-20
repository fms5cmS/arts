package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: head}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			tmp := cur.Next
			cur.Next = cur.Next.Next
			tmp.Next = nil
		} else {
			// 注意，只有在 else 的时候才移动 cur！否则如果连续多个节点的值均为 val 时，会出现错误
			// 如 [7, 7, 7, 7] val = 7 时
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		array  []int
		val    int
		output []int
	}{
		{
			array:  []int{1, 2, 6, 3, 4, 5, 6},
			val:    6,
			output: []int{1, 2, 3, 4, 5},
		},
		{
			array:  []int{},
			val:    1,
			output: []int{},
		},
		{
			array:  []int{7, 7, 7, 7},
			val:    7,
			output: []int{},
		},
	}
	for _, test := range tests {
		before := generateListViaArray(test.array)
		after := removeElements(before, test.val)
		assert.Equal(t, test.output, generateArrayViaList(after))
	}
}
