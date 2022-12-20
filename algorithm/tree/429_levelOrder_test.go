package tree

func levelOrder429(root *Node) [][]int {
	result := make([][]int, 0)
	queue := make([]*Node, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur.Val)
			queue = append(queue, cur.Children...)
		}
		result = append(result, tmp)
	}
	return result
}
