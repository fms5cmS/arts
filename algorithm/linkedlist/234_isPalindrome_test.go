package linkedlist

import (
	"fmt"
	"testing"
)

// 1 -> 2 -> 2 -> 1
// 回文串是对称的，所以正着读和反着读应该一致
// 1. 遍历链表将其值依次入栈，然后再次遍历链表，并将值与出栈元素对比
func isPalindromeStack(head *ListNode) bool {
	stack := make([]*ListNode, 0)
	cur := head
	// 入栈
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}
	// 对比
	cur = head
	for i := len(stack) - 1; i >= 0; i-- {
		tmp := stack[i]
		if tmp.Val != cur.Val {
			return false
		}
		cur = cur.Next
	}
	return true
}

// 快慢指针找到链表的中点，然后将中点后的节点翻转，再从左往右比较即可
// 注意：会破坏链表的结构
func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	// 1—>2—>3—>4—>5  slow 是值为 3 的节点
	// 1—>2—>3—>4     slow 是值为 2 的节点
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := reverseLinkedListForPalindrome(nil, slow.Next)
	left := head
	// 这里是为了最后复原链表结构的
	last := right
	// 判断是否为回文结构
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next
	}

	// 这里实际链表的结构已经被改变了，如何重新复原链表的结构呢
	slow.Next = reverseLinkedListForPalindrome(nil, last)
	return true
}

// pre 代表了 cur 翻转后指向的节点
// 返回翻转后链表的头节点
func reverseLinkedListForPalindrome(pre, cur *ListNode) *ListNode {
	for cur != nil {
		temp := cur.Next     // 防止链表断开，所以需要一个临时变量来保存
		cur.Next = pre       // 翻转 cur 节点，使其 Next 指向 pre
		pre, cur = cur, temp // 更新 pre，cur 以便翻转剩下的节点
	}
	return pre
}

func TestIsPalindrome(t *testing.T) {
	headOdd := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}
	headEven := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	fmt.Println(isPalindrome(headOdd))
	fmt.Println(isPalindrome(headEven))
	fmt.Println(headOdd)
	fmt.Println(headEven)
}
