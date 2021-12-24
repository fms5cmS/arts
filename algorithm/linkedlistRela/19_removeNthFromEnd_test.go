package linkedlistRela

import (
	"fmt"
	"testing"
)

// 假设链表长度为 N，倒数第 n 个节点的索引其实为 N-n！N 个节点从 head 到 tail 共需要 N-1 步。
// 1. 快指针从 head 先出发，每次走一步，先走 n 步；
//   - PS：此时快指针距离走到尾节点还有 (N-1)-n 步
// 2. 慢指针从 head 后出发，和快指针一起每次走一步，直到快指针走到最后一个节点；
//   - PS：此时慢指针走了 (N-1)-n 步，索引为 (N-1)-n，正好是 N-n 的前一位
// 3. 此时慢指针所在位置为要删除节点的前一个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	slow, fast := dummy, dummy
	for n > 0 && fast != nil {
		fast = fast.Next
		n--
	}
	fast = fast.Next
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 原理同上
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	pre, cur := dummy, head
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			pre = pre.Next
		}
		i++
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}

func TestRemoveNthFromEnd(t *testing.T) {
	list := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 5, Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	a := removeNthFromEnd(list, 7)
	fmt.Println(a.String())
}
