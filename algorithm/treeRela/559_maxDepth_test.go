package treeRela

// 与 104 的思路是一致的
func maxDepthForNTree(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for _, child := range root.Children {
		depth = max(depth, maxDepthForNTree(child))
	}
	return depth+1
}
