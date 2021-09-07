package treeRela

import "testing"

// 给定的是完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点
func connect116(root *ConnectNode) *ConnectNode {
	queue := make([]*ConnectNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			// i < size-1 说明 cur 并不是本层的最后一个节点，queue 中还有其他的节点
			// 而本层的最后一个节点(i == size-1) 的 Next 指向 nil
			if i < size-1 {
				cur.Next = queue[0]
			}
		}
	}
	return root
}

// 层序遍历，每一层的节点串起来
func connect(root *ConnectNode) *ConnectNode {
	if root == nil {
		return nil
	}
	curLevel, nextLevel := make([]*ConnectNode, 0), make([]*ConnectNode, 0, 1)
	nextLevel = append(nextLevel, root)
	level := 1
	for len(nextLevel) != 0 {
		curLevel = make([]*ConnectNode, 0)
		curLevel = append(curLevel, nextLevel...)
		nextLevel = make([]*ConnectNode, 0, 1<<level)
		level++
		for i := 0; i < len(curLevel); i++ {
			if i+1 < len(curLevel) {
				curLevel[i].Next = curLevel[i+1]
			}
			if curLevel[i].Left != nil {
				nextLevel = append(nextLevel, curLevel[i].Left)
			}
			if curLevel[i].Right != nil {
				nextLevel = append(nextLevel, curLevel[i].Right)
			}
		}
	}
	return root
}

func TestConnect(t *testing.T) {
	root := &ConnectNode{
		Val:   1,
		Left:  &ConnectNode{Val: 2, Left: &ConnectNode{Val: 4}, Right: &ConnectNode{Val: 5}},
		Right: &ConnectNode{Val: 3, Left: &ConnectNode{Val: 6}, Right: &ConnectNode{Val: 7}},
	}
	connect(root)
}

// 递归来连接每一层的节点
//         1
//    2          3
// 4     5    6      7
func connect2(root *ConnectNode) *ConnectNode {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

// 假设传入的是 2、3
func connectTwoNode(first, second *ConnectNode) {
	if first == nil || second == nil {
		return
	}
	first.Next = second
	// 连接 first、second 各自的子节点（4、5 连接，6、7 连接）
	connectTwoNode(first.Left, first.Right)
	connectTwoNode(second.Left, second.Right)
	// 连接跨越父节点的两个节点（5、6 连接）
	connectTwoNode(first.Right, second.Left)
}
