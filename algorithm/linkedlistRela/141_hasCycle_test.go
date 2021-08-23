package linkedlistRela

// 使用快慢指针，如果有环的话，两个指针一定会相遇！
// slow 每次一步，fast 每次两步
// 如果有环的话，fast 一定先进入环中，如果 fast 和 slow 相遇的话，一定是在环中相遇
func hasCycle(head *ListNode) bool {
	if head == nil{
		return false
	}
	// 这里注意不能都是 head，否则 for 循环里直接返回 true 了
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// 每次遍历一个节点，先从 Set 中找有没有经过该节点，有的话说明链表有环，返回 true
// 否则就将该节点保存到 Set 中
func hasCycleBySet(head *ListNode) bool {
	mem := make(map[*ListNode]int)
	for head != nil {
		if v := mem[head]; v != 0 {
			return true
		}
		mem[head] = 1
		head = head.Next
	}
	return false
}
