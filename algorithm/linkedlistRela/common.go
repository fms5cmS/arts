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
