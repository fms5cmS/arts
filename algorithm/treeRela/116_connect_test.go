package treeRela

import "testing"

// 层序遍历，每一层的节点串起来
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	curLevel, nextLevel := make([]*Node, 0), make([]*Node, 0, 1)
	nextLevel = append(nextLevel, root)
	level := 1
	for len(nextLevel) != 0 {
		curLevel = make([]*Node, 0)
		curLevel = append(curLevel, nextLevel...)
		nextLevel = make([]*Node, 0, 1<<level)
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
	root := &Node{
		Val:   1,
		Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
		Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}},
	}
	connect(root)
}

// 递归来连接每一层的节点
//         1
//    2          3
// 4     5    6      7
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

// 假设传入的是 2、3
func connectTwoNode(first, second *Node) {
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
