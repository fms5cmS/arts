package treeRela

import (
	"fmt"
	"testing"
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	if root == nil {
		return result
	}
	var traversal func(cur *TreeNode, count int, path []int)
	traversal = func(cur *TreeNode, count int, path []int) {
		// 遇到叶子节点且找到和为 targetSum 的路径
		if cur.Left == nil && cur.Right == nil && count == 0 {
			result = append(result, path)
			return
		}
		// 遇到叶子节点，但和不满足 targetSum
		if cur.Left == nil && cur.Right == nil {
			return
		}
		if cur.Left != nil {
			pathTmp := make([]int, 0, len(path)+1)
			pathTmp = append(pathTmp, path...)
			pathTmp = append(pathTmp, cur.Left.Val)
			count -= cur.Left.Val
			traversal(cur.Left, count, pathTmp)
			// 回溯
			count += cur.Left.Val
		}
		if cur.Right != nil {
			pathTmp := make([]int, 0, len(path)+1)
			pathTmp = append(pathTmp, path...)
			pathTmp = append(pathTmp, cur.Right.Val)
			count -= cur.Right.Val
			traversal(cur.Right, count, pathTmp)
			// 回溯
			count += cur.Right.Val
		}
	}
	path = append(path, root.Val)
	traversal(root, targetSum-root.Val, path)
	return result
}

//                         5
//                4                  8
//          11                  13       4
//       7      2                      5    1
func TestPathSum(t *testing.T) {
	root := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}}
	result := pathSum(root, 22)
	for _, path := range result {
		fmt.Println(path)
	}
}
