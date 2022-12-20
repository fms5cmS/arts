package tree

import "fmt"

// 找一个树里重复的子树
// 重点：比较子树 ————> 通过将子树序列化为字符串后，比较字符串来完成
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	strMap := make(map[string]int)
	ret := make([]*TreeNode, 0)
	// 定义一个函数来将树序列化（这里是通过中序遍历来实现的，也可以用其他遍历的方式），以便于比较树
	var marshalTreeByInOrder func(root *TreeNode) string
	marshalTreeByInOrder = func(root *TreeNode) string {
		if root == nil {
			return "*"
		}
		leftStr := marshalTreeByInOrder(root.Left)
		rightStr := marshalTreeByInOrder(root.Right)
		str := fmt.Sprintf("%d_%s_%s", root.Val, leftStr, rightStr)
		// 这里使用了闭包，不用再将 map、ret 作为参数传入了
		// 该树的数量自增
		strMap[str]++
		// 如果某棵树的数量为 2，则将其加入返回结果中
		// 注意：这里只有等于 2 的时候才记录，如果是 num > 1 的话会导致相同结构的子树重复放入
		if num := strMap[str]; num == 2 {
			ret = append(ret, root)
		}
		return str
	}
	marshalTreeByInOrder(root)
	return ret
}
