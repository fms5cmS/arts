package treeRela

import (
	"math"
	"testing"
)

// 注意题目说的是二叉搜索树，而二叉搜索树是有序的！
// 所以可以当作是在一个有序数组中求解两个数的最小差值
// 最直观的想法，就是把二叉搜索树转换成有序数组，然后遍历一遍数组，就统计出来最小差值了。
func getMinimumDifference(root *TreeNode) int {
	var res []int
	findMIn(root, &res)
	min := math.MaxInt64 //一个比较大的值
	for i := 1; i < len(res); i++ {
		tempValue := res[i] - res[i-1]
		if tempValue < min {
			min = tempValue
		}
	}
	return min
}

// 中序遍历，里面会 append 操作，所以切片的传入要使用指针，防止底层数组改变
func findMIn(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	findMIn(root.Left, res)
	*res = append(*res, root.Val)
	findMIn(root.Right, res)
}

func TestGetMinimumDifference(t *testing.T) {
	root := ConstructTreeByLevelOrder([]interface{}{1, 0, 48, nil, nil, 12, 49})
	t.Log(getMinimumDifference(root))
}
