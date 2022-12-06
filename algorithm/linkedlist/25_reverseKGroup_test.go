package linkedlist

import (
	"fmt"
	"testing"
)

// 每 k 个节点一组进行翻转，请你返回翻转后的链表
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, boundary := dummy, dummy
	// 下面需要分割链表，boundary 为初始分割点
	for boundary != nil {
		// 以 k 为长度，分割链表，找到最新的 boundary
		for i := 0; i < k && boundary != nil; i++ {
			boundary = boundary.Next
		}
		if boundary == nil {
			break
		}
		// 分割后会得到两个链表，这里分别记录两个链表的头部
		firstStart, nextStart := pre.Next, boundary.Next
		// 断开两个链表的关联关系，便于后面翻转单个链表
		boundary.Next = nil
		// 翻转第一个链表，并使其尾节点指向第二个链表的头部
		// 实际翻转的是 "初始 boundary ~ 最新 boundary" 的这个链表
		fmt.Println(firstStart)
		pre.Next = reverseListWithPre(nextStart, firstStart)
		// 此时 firstStart 为翻转后链表的尾节点
		pre = firstStart
		boundary = firstStart // 重置初始分割点
	}
	return dummy.Next
}

// 翻转以 head 为头节点的链表，翻转后链表的尾节点指向 pre
// 返回翻转后的头节点
func reverseListWithPre(pre, head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre, cur = cur, tmp
	}
	return pre
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		k    int
		nums []int
		want []int
	}{
		{
			name: "1",
			k:    2,
			nums: []int{1, 2, 3, 4, 5},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "2",
			k:    3,
			nums: []int{1, 2, 3, 4, 5},
			want: []int{3, 2, 1, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			originalList := generateListViaArray(test.nums)
			t.Logf("%s, original list: %s", test.name, originalList)
			reversedList := reverseKGroup(originalList, test.k)
			t.Logf("%s, reversed list: %s, with %d-group", test.name, reversedList, test.k)
		})
	}
}
