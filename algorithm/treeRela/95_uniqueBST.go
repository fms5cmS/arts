package treeRela

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateBetween(1, n)
}

func generateBetween(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := make([]*TreeNode, 0)
	// 枚举可行的根节点
	for i := start; i <= end; i++ {
		// 获取所有可行的左右子树集合
		leftTrees := generateBetween(start, i-1)
		rightTrees := generateBetween(i+1, end)
		// 从左右子树集合各选出一个子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				curTree := &TreeNode{i, nil, nil}
				curTree.Left = left
				curTree.Right = right
				allTrees = append(allTrees, curTree)
			}
		}
	}
	return allTrees
}
