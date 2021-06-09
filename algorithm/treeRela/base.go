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
	// nextQueue 保存下一层的所有节点，curQueue 保存当前遍历的节点
	nextQueue, curQueue := make([]*TreeNode, 0), make([]*TreeNode, 0)
	ret := make([][]int, 0)
	// 1. 往 nextQueue 中加入根节点
	nextQueue = append(nextQueue, root)
	// i 用于记录遍历的第几层，对应二维数组 ret 的行数
	i := 0
	// 2. 遍历，知道下一层没有节点
	for len(nextQueue) != 0 {
		// 获得当前要遍历的那一层所有节点
		curQueue = nextQueue
		// nextQueue 清空
		nextQueue = make([]*TreeNode, 0)
		// 初始化 ret 的数组
		tmpArr := make([]int, 0)
		ret = append(ret, tmpArr)
		// 开始遍历当前层
		for _, node := range curQueue {
			// 当前层的值填入 ret 中
			ret[i] = append(ret[i], node.Val)
			// 向 nextQueue 中放入将要遍历的下一层节点
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		// 当前行遍历完，i++
		i++
	}
	return ret
}
