package linkedlistRela

import "testing"

// 迭代解法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = nil
	for head != nil {
		tmp := head.Next
		// head.Next, pre = pre, head
		head = tmp
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
	list := new(ListNode)
	list.Val = 1
	list.Next = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}
	t.Log(list)
	// node := reverseListRecursion(list)
	node := reverseList(list)
	t.Log(node)
}
