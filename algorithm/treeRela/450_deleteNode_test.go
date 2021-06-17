package treeRela

func deleteNodeDetail(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// 查找 Val 等于 key 的节点
	if key > root.Val {
		root.Right = deleteNodeDetail(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNodeDetail(root.Left, key)
	} else {
		// 找到后，执行删除操作
		// 如果该节点没有子节点，将自己删除
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 经过上面的条件后，这里至少有一个子节点不为空
		// 如果该节点只有一个子节点，子节点替换父节点
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		// 如果该节点有两个子节点：
		// 方法一：用其右子树的最小值替换自身(这里选择的是该方法)
		// 方法二：用其左子树的最大值替换自身
		if root.Left != nil && root.Right != nil {
			minNode := getMinNodeOfBST(root.Right)
			root.Val = minNode.Val
			root.Right = deleteNodeDetail(root.Right, minNode.Val)
		}
	}
	return root
}

// 获取 BST 最小值的节点
func getMinNodeOfBST(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// 简化部分重复逻辑
func deleteNodeSimple(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// 查找 Val 等于 key 的节点
	if key > root.Val {
		root.Right = deleteNodeSimple(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNodeSimple(root.Left, key)
	} else {
		// 找到后，执行删除操作
		// 这两个 if 可以同时处理 root 没有、只有一个子节点的情况
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 如果该节点有两个子节点：
		// 方法一：用其右子树的最小值替换自身(这里选择的是该方法)
		// 方法二：用其左子树的最大值替换自身
		minNode := getMinNodeOfBST(root.Right)
		// 注意：实际中并不会通过修改节点值来交换节点，因为实际的节点中数据结构会比较复杂！交换的方法见下面
		root.Val = minNode.Val
		root.Right = deleteNodeSimple(root.Right, minNode.Val)
	}
	return root
}


// 交换节点，而不再是交换节点的值
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// 查找 Val 等于 key 的节点
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		// 找到后，执行删除操作
		// 这两个 if 可以同时处理 root 没有、只有一个子节点的情况
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 如果该节点有两个子节点：
		// 方法一：用其右子树的最小值替换自身(这里选择的是该方法)
		// 方法二：用其左子树的最大值替换自身
		// 交换节点，而不再是交换节点的值！
		node := root.Right
		for node.Left != nil {
			node = node.Left
		}
		node.Left, root = root.Left, root.Right
	}
	return root
}
