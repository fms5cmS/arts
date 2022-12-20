package tree

import (
	"fmt"
	"math"
	"testing"
)

// 如果当前节点要做的事需要通过左右子树的计算结果，就要用到后序遍历
func maxSumBST(root *TreeNode) int {
	maxSum := 0
	traverseForMaxSum(root, &maxSum)
	return maxSum
}

func traverseForMaxSum(root *TreeNode, maxSum *int) (isBST bool, min, max, sum int) {
	if root == nil {
		return true, math.MaxInt64, math.MinInt64, 0
	}
	// 后序遍历
	leftIsBST, leftMin, leftMax, leftSum := traverseForMaxSum(root.Left, maxSum)
	rightIsBST, rightMin, rightMax, rightSum := traverseForMaxSum(root.Right, maxSum)
	var rootIsBST bool
	var rootMin, rootMax, rootSum int
	if leftIsBST && rightIsBST && root.Val > leftMax && root.Val < rightMin {
		// root 是 BST
		rootIsBST = true
		rootMin = compareTwoInt(leftMin, root.Val, true)
		rootMax = compareTwoInt(root.Val, rightMax, false)
		rootSum = leftSum + rightSum + root.Val
		*maxSum = compareTwoInt(rootSum, *maxSum, false)
	} else {
		// root 不是 BST，其他字段不会用到，没必要计算
		rootIsBST = false
	}
	return rootIsBST, rootMin, rootMax, rootSum
}

func compareTwoInt(x, y int, min bool) int {
	if min {
		if x < y {
			return x
		}
		return y
	}
	if x > y {
		return x
	}
	return y
}

func TestMaxSumBST(t *testing.T) {
	root := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}}
	fmt.Println(maxSumBST(root))
}
