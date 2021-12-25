package treeRela

// n 叉树的后序遍历
// 类同 145
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	result := make([]int, 0)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		for i, child := range cur.Children {
			if child != nil {
				stack = append(stack, cur.Children[i])
			}
		}
	}
	reverseIntsFor590(result)
	return result
}

func reverseIntsFor590(src []int) {
	left, right := 0, len(src)-1
	for left < right {
		src[left], src[right] = src[right], src[left]
		left++
		right--
	}
}
