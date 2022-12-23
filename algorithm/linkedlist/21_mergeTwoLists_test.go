package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "1",
			nums1: []int{1, 2, 4},
			nums2: []int{1, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:  "2",
			nums1: []int{},
			nums2: []int{},
			want:  []int{},
		},
		{
			name:  "3",
			nums1: []int{},
			nums2: []int{0},
			want:  []int{0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list1 := generateListViaArray(test.nums1)
			list2 := generateListViaArray(test.nums2)
			l1 := mergeTwoLists(list1, list2)
			get := generateArrayViaList(l1)
			assert.Equal(t, test.want, get)

			list11 := generateListViaArray(test.nums1)
			list22 := generateListViaArray(test.nums2)
			l2 := mergeTwoLists(list11, list22)
			get22 := generateArrayViaList(l2)
			assert.Equal(t, test.want, get22)
		})
	}
}
