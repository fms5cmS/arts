package treeRela

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder 层序遍历，将其放入而未数组中
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// nextLevel 保存下一层的所有节点，curLevel 保存当前遍历的节点
	nextLevel, curLevel := make([]*TreeNode, 0), make([]*TreeNode, 0)
	ret := make([][]int, 0)
	// 1. 往 nextLevel 中加入根节点
	nextLevel = append(nextLevel, root)
	// i 用于记录遍历的第几层，对应二维数组 ret 的行数
	i := 0
	// 2. 遍历，直到下一层没有节点
	for len(nextLevel) != 0 {
		// 获得当前要遍历的那一层所有节点
		curLevel = nextLevel
		// nextLevel 清空
		nextLevel = make([]*TreeNode, 0)
		// 初始化 ret 的数组
		tmpArr := make([]int, 0)
		ret = append(ret, tmpArr)
		// 开始遍历当前层
		for _, node := range curLevel {
			// 当前层的值填入 ret 中
			ret[i] = append(ret[i], node.Val)
			// 向 nextLevel 中放入将要遍历的下一层节点
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		// 当前行遍历完，i++
		i++
	}
	return ret
}

type ConnectNode struct {
	Val   int
	Left  *ConnectNode
	Right *ConnectNode
	Next  *ConnectNode
}

func getIndex(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

// N叉树的节点
type Node struct {
	Val      int
	Children []*Node
}

func constructTreeByArray(array []interface{}) *TreeNode {
	if len(array) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(array))
	for i, v := range array {
		if v != nil {
			nodes[i] = &TreeNode{Val: v.(int)}
		} else {
			nodes[i] = nil
		}
	}
	// 由于 array 并不是按照二叉树顺序存储的结构来放置的，所以不能用顺序存储的方式来生成树结构
	// 记录子节点移动的位置
	j := 1
	for i := 0; i < len(array); i++ {
		if nodes[i] != nil {
			if j < len(array) {
				nodes[i].Left = nodes[j]
				j++
			}
			if j < len(array) {
				nodes[i].Right = nodes[j]
				j++
			}
		}
	}
	return nodes[0]
}
