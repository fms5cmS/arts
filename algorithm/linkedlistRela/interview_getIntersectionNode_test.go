package linkedlistRela

// 面试题 02.07. 链表相交
// ListNode 只有一个 Next 指针指向下一个节点，所以当两个链表相交后，之后的长度是一样的！
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	lenA, lenB := 0, 0
	// 求两个链表的长度
	for curA != nil {
		lenA++
		curA = curA.Next
	}
	for curB != nil {
		lenB++
		curB = curB.Next
	}
	// 获取两个链表的长度差 step
	step := 0
	var more, less *ListNode
	if lenA > lenB {
		step = lenA - lenB
		more, less = headA, headB
	} else {
		step = lenB - lenA
		more, less = headB, headA
	}
	// 较长的链表，指针先走 step 步
	for i := 0; i < step; i++ {
		more = more.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for more != less {
		more, less = more.Next, less.Next
	}
	return more
}
