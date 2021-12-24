package treeRela

import (
	"fmt"
	"strconv"
	"strings"
)

// 因为是从根节点到叶子节点的路径，需要父节点指向子节点，所以必须是**前序遍历**
// 实现还需要回溯来回退一个路径再进入另一个路径
func binaryTreePaths(root *TreeNode) []string {
	path, result := make([]int, 0), make([]string, 0)
	if root == nil {
		return result
	}
	traversal(root, &path, &result)
	return result
}

// 递归 + 回溯
func traversal(cur *TreeNode, path *[]int, result *[]string) {
	*path = append(*path, cur.Val)
	// 叶子节点
	if cur.Left == nil && cur.Right == nil {
		sPath := strings.Builder{}
		for i := 0; i < len(*path)-1; i++ {
			sPath.WriteString(fmt.Sprintf("%d->", (*path)[i]))
		}
		sPath.WriteString(strconv.Itoa((*path)[len(*path)-1]))
		*result = append(*result, sPath.String())
	}
	if cur.Left != nil {
		traversal(cur.Left, path, result)
		*path = (*path)[:len(*path)-1] // 回溯
	}
	if cur.Right != nil {
		traversal(cur.Right, path, result)
		*path = (*path)[:len(*path)-1] // 回溯
	}
}

func binaryTreePaths2(root *TreeNode) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}
	var travel func(cur *TreeNode, srcPath string)
	travel = func(cur *TreeNode, srcPath string) {
		if cur.Left == nil && cur.Right == nil {
			result = append(result, fmt.Sprintf("%s->%d", srcPath, cur.Val))
			return
		}
		// 这两步其实也是有回溯的，每次 travel 调用之后，srcPath 依然和原本的值一致，相当于回溯了
		if cur.Left != nil {
			travel(cur.Left, fmt.Sprintf("%s->%d", srcPath, cur.Left.Val))
		}
		if cur.Right != nil {
			travel(cur.Right, fmt.Sprintf("%s->%d", srcPath, cur.Right.Val))
		}
	}
	travel(root, strconv.Itoa(root.Val))
	return result
}
