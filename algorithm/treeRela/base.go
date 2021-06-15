package treeRela

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历，将其放入而未数组中
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

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func getIndex(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}
