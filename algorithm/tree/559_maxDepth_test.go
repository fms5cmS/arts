package tree

import "arts/algorithm/utils"

// 与 104 的思路是一致的
func maxDepthForNTree(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	for _, child := range root.Children {
		depth = utils.Max(depth, maxDepthForNTree(child))
	}
	return depth + 1
}
