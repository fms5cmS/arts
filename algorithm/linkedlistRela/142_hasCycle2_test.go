package linkedlistRela

func detectCycle(head *ListNode) *ListNode {
	// 链表长度为 0 或 1 时无环
	if head == nil || head.Next == nil {
		return nil
	}
	// 快、慢指针一起走，快指针每次走两步，慢指针每次走一步
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		// 快慢指针相遇
		if slow == fast {
			// head 与慢指针同时从各自当前位置开始走，且速度相同
			for head != slow {
				head, slow = head.Next, slow.Next
			}
			// 当 head 与慢指针相遇时，该节点即为入环节点
			return head
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
