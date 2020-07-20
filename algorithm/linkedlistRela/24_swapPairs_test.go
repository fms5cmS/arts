package linkedlistRela

import (
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
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	t.Log(list)
	t.Log(swapPairs(list))
}
