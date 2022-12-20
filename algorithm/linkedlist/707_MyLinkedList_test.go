package linkedlist

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type MyLinkedList struct {
	dummyHead *MyListNode
	size      int
}

func (this *MyLinkedList) String() string {
	str := strings.Builder{}
	cur := this.dummyHead
	for cur.Next != nil {
		str.WriteString(strconv.Itoa(cur.Next.Val))
		str.WriteString(" --> ")
		cur = cur.Next
	}
	str.WriteString("nil")
	return str.String()
}

type MyListNode struct {
	Val  int
	Next *MyListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{dummyHead: &MyListNode{Val: 0}}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.dummyHead.Next
	for index > 0 {
		cur = cur.Next
		index--
	}
	return cur.Val
}

// Add a node of value val before the first element of the linked list. After the insertion,
// the new node will be the first node of the linked list.
func (this *MyLinkedList) AddAtHead(val int) {
	beforeHead := this.dummyHead.Next
	newHead := &MyListNode{Val: val, Next: beforeHead}
	this.dummyHead.Next = newHead
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &MyListNode{Val: val}
	this.size++
}

// Add a node of value val before the index-th node in the linked list.
// If index equals to the length of linked list, the node will be appended to the end of linked list.
// If index is greater than the length, the node will not be inserted.
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.size {
		this.AddAtTail(val)
		return
	} else if index < 0 {
		this.AddAtHead(val)
		return
	} else if index > this.size {
		return
	}
	pre := this.dummyHead
	for index > 0 {
		pre = pre.Next
		index--
	}
	tmp := pre.Next
	pre.Next = &MyListNode{Val: val, Next: tmp}
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	pre := this.dummyHead
	for index > 0 {
		pre = pre.Next
		index--
	}
	tmp := pre.Next
	pre.Next = pre.Next.Next
	tmp.Next = nil
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func TestMyLinkedList(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	fmt.Println(list.String()) // 1
	list.AddAtTail(3)
	fmt.Println(list.String()) //  1 -> 3
	list.AddAtIndex(1, 2)
	fmt.Println(list.String()) // 1 -> 2 -> 3
	fmt.Println(list.Get(1))   // 2
	list.DeleteAtIndex(1)
	fmt.Println(list.String()) // 1 - > 3
	fmt.Println(list.Get(1))   // 3
}

func TestMyLinkedList2(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	fmt.Println(list.String())
	list.DeleteAtIndex(0)
	fmt.Println(list.String())
}
