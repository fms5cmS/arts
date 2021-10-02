package treeRela

// 注意：不能直接每次遍历右节点，因为如果右节点不存在而左节点的话，树的右视图是可以看到左节点的
// 使用层序遍历，如果遍历到了当前层最后一个节点，就将其添加到结果集中
func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if i == size-1 {
				result = append(result, cur.Val)
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return result
}
