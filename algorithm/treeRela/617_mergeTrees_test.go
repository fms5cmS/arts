package treeRela

// 如何同时遍历两个二叉树呢？和遍历一个树逻辑是一样的，只不过传入两个树的节点，同时操作。
// 递归，前序遍历（中、后序也可以）
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root := new(TreeNode)
	root.Val = root1.Val + root2.Val
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

// 迭代
// 思路同 101，把两棵树的节点都入队再处理
func mergeTreesNotRecursion(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue := []*TreeNode{root1, root2}
	for len(queue) > 0 {
		node1, node2 := queue[0], queue[1]
		queue = queue[2:]
		// 两个节点一定不为空，val 相加
		node1.Val += node2.Val
		// 两棵树的左右节点不为空，入队
		if node1.Left != nil && node2.Left != nil {
			queue = append(queue, node1.Left, node2.Left)
		}
		if node1.Right != nil && node2.Right != nil {
			queue = append(queue, node1.Right, node2.Right)
		}
		// 注意：这里没有入队操作
		// node1 的左节点为空，node2 的不为空，将 node2 的值赋值过去
		if node1.Left == nil && node2.Left != nil {
			node1.Left = node2.Left
		}
		// node1 的右节点为空，node2 的不为空，将 node2 的值赋值过去
		if node1.Right == nil && node2.Right != nil {
			node1.Right = node2.Right
		}
	}
	return root1
}
