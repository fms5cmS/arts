package linkedlist

import "testing"

// 每 k 个节点一组进行翻转，请你返回翻转后的链表
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre, boundary := dummy, dummy
	for boundary.Next != nil {
		// boundary 用于分组，每 k 个节点为一组
		for i := 0; i < k && boundary != nil; i++ {
			boundary = boundary.Next
		}
		if boundary == nil {
			break
		}
		// 分别代表了两个分组各自的头节点
		start, next := pre.Next, boundary.Next
		// 酱两个分组断开连接，便于下一步的翻转
		boundary.Next = nil
		// 翻转第一个分组，翻转完成后， start.Next 会指向 next
		pre.Next = reverseOfList(next, start)
		// 重置 pre、boundary 的位置
		pre, boundary = start, start
	}
	return dummy.Next
}

// 翻转以 cur 为起点的链表，翻转后的尾节点指向 pre
func reverseOfList(pre, cur *ListNode) *ListNode {
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre, cur = cur, tmp
	}
	return pre
}

func TestReverseKGroup(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
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
		},
	}
	t.Log(list)
	t.Log(reverseKGroup(list, 3))
}
