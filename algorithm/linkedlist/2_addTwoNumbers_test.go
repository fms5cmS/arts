package linkedlist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	var val1, val2, now, carry int
	for l1 != nil || l2 != nil {
		// 这里每次必须要重置 val1 和 val2 的值，每次循环的默认值为 0！
		val1, val2 = 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next // 由于 l1 当前节点的值已被放入 val1，所以这里可以提前
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		now, carry = getNowAndCarry(val1+val2, carry)
		cur.Next = &ListNode{Val: now}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}

func getNowAndCarry(x, y int) (now, carry int) {
	val := x + y
	now, carry = val%10, val/10
	return
}
