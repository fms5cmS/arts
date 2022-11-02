package linkedlistRela

// ListNode 只有一个 Next 指针指向下一个节点，所以当两个链表相交后，之后的长度是一样的！
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 获取两个链表的长度
	stepA, stepB := 0, 0
	curA, curB := headA, headB
	for ; curA != nil; stepA++ {
		curA = curA.Next
	}
	for ; curB != nil; stepB++ {
		curB = curB.Next
	}
	// 计算两个链表的长度差
	div := stepA - stepB
	// 区分长短链表，fast 先在长的链表上移动 div 步
	var slow, fast *ListNode
	if div < 0 {
		fast, slow = headB, headA
		div = -div
	} else {
		fast, slow = headA, headB
	}
	for i := 0; i < div; i++ {
		fast = fast.Next
	}
	// 如果两个链表相交，则必定会在后面的移动中指针相等
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
