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

func ConstructTreeByLevelOrder(levelVals []interface{}) *TreeNode {
	if len(levelVals) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(levelVals))
	var root *TreeNode
	// 1. 数值数组转为节点数组（线性存储结构）
	for i, val := range levelVals {
		if val != nil {
			nodes[i] = &TreeNode{Val: val.(int)}
		} else {
			nodes[i] = nil
		}
	}
	// 最开始已经对空数组做了处理，所以走到这里的话 nodes 的长度一定是大于 0 的！
	root = nodes[0]
	// 2. 拼接树：线性存储结构转链式存储结构，因为最大要用 i*2+2 的索引下标，所以需要以此作为判断条件防止下标越界
	for i := 0; i*2+2 < len(levelVals); i++ {
		if nodes[i] != nil {
			nodes[i].Left = nodes[i*2+1]
			nodes[i].Right = nodes[i*2+2]
		}
	}
	return root
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

func getIndex(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}
