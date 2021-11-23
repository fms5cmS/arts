package linkedlistRela

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	start, end := 0, len(nodes)-1
	for start < end {
		nodes[start].Next = nodes[end]
		start++
		if start == end {
			break
		}
		nodes[end].Next = nodes[start]
		end--
	}
	nodes[start].Next = nil
}
