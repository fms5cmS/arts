package linkedlist

// 1 -> 2 -> 3 -> 4 -> 5, left = 2, right = 4
// 1 -> 4 -> 3 -> 2 -> 5
// cur1 指向 1，cur2 指向 3
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 使用虚拟头节点
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	// 从虚拟头节点开始走 left-1 步，走到 left 的前一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	// 从 pre 开始走 right-left+1 步，走到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	// 截取链表
	leftNode := pre.Next
	cur := rightNode.Next
	// 切断连接
	pre.Next, rightNode.Next = nil, nil
	// 翻转链表的子区间
	reverseLinkedList(leftNode)
	// 将翻转后的链表再次接回去
	pre.Next = rightNode
	leftNode.Next = cur
	return dummyNode.Next
}

func reverseLinkedList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

// 只遍历一遍链表，头插法
// 9 -> 7 -> 2 -> 5 -> 4 -> 3 -> 6，翻转 从值 2 到 3 之间的节点（2，5，4，3）
// 9 -> 7 -> 5 -> 2 -> 4 -> 3 -> 6   5 插入到 2 前面
// 9 -> 7 -> 4 -> 5 -> 2 -> 3 -> 6   4 插入到 5 前面
// 9 -> 7 -> 3 -> 4 -> 5 -> 2 -> 6   3 插入到 4 前面
func reverseBetweenOnce(head *ListNode, left int, right int) *ListNode {
	// 使用虚拟头节点
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	// 从虚拟头节点开始走 left-1 步，走到 left 的前一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
