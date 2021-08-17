package linkedlistRela

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			tmp := cur.Next
			cur.Next = cur.Next.Next
			tmp.Next = nil
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
