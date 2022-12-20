package linkedlist

// 使用快慢指针，如果有环的话，两个指针一定会相遇！
// slow 每次一步，fast 每次两步
// 如果有环的话，fast 一定先进入环中，如果 fast 和 slow 相遇的话，一定是在环中相遇
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
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
