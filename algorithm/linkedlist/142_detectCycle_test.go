package linkedlist

// 数学证明：https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0142.%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8II.md
func detectCycle(head *ListNode) *ListNode {
	// 链表长度为 0 或 1 时无环
	if head == nil || head.Next == nil {
		return nil
	}
	// 快、慢指针一起走，快指针每次走两步，慢指针每次走一步
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		// 快慢指针相遇（不一定在环的起点），说明有环
		if slow == fast {
			// 重置 slow 为 head，然后slow 与 fast 同时从各自当前位置开始走，且速度相同
			slow = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}
			// 当 slow 与 fast 相遇时，该节点即为入环节点
			return slow
		}
	}
	return nil
}

// 通过 Set 保存每个节点
// 当第一次遇到该节点已保存在 Set 的情况时，此节点就是入环的第一个节点
func detectCycle_set(head *ListNode) *ListNode {
	mem := make(map[*ListNode]int)
	for head != nil {
		if mem[head] == 0 {
			mem[head] = 1
		} else {
			return head
		}
		head = head.Next
	}
	return nil
}
