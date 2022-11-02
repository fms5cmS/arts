package linkedlistRela

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	cur := node
	str := strings.Builder{}
	for cur.Next != nil {
		str.WriteString(strconv.Itoa(cur.Val))
		str.WriteString(" -> ")
		cur = cur.Next
	}
	str.WriteString(strconv.Itoa(cur.Val))
	return str.String()
}

func generateListViaArray(array []int) *ListNode {
	if len(array) < 1 {
		return nil
	}
	head := &ListNode{Val: array[0]}
	cur := head
	for i := 1; i < len(array); i++ {
		cur.Next = &ListNode{Val: array[i]}
		cur = cur.Next
	}
	return head
}

func generateArrayViaList(head *ListNode) []int {
	array := make([]int, 0)
	cur := head
	for ; cur != nil; cur = cur.Next {
		array = append(array, cur.Val)
	}
	return array
}
