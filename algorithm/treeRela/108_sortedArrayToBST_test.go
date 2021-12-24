package treeRela

import (
	"fmt"
	"testing"
)

// 先做：106、654、701、450
// 有序数组转为平衡二叉搜索树
// 数组构造二叉树：本质就是寻找分割点，分割点作为当前节点，然后递归左区间和右区间。
// 有序数组的分割点就是数组中间位置的元素，如果数组长度为偶数，两个中间节点随便取即可，只不过构成的平衡二叉搜索树不同而已，所以同一个数组答案不唯一
func sortedArrayToBST(nums []int) *TreeNode {
	var traversal func(num []int, left, right int) *TreeNode
	traversal = func(num []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)>>1
		root := &TreeNode{Val: num[mid]}
		root.Left = traversal(num, left, mid-1)
		root.Right = traversal(num, mid+1, right)
		return root
	}
	// 注意这里传参的 0、len(nums)-1，可以看 704 的左闭右闭区间使用
	return traversal(nums, 0, len(nums)-1)
}

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	tree := sortedArrayToBST(nums)
	fmt.Println(tree.Val)
}
