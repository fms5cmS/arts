package treeRela

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index, max := maxOfNums(nums)
	root := new(TreeNode)
	root.Val = max
	// 左闭右开区间，所以这里是 index
	root.Left = constructMaximumBinaryTree(nums[:index])
	if index < len(nums) {
		root.Right = constructMaximumBinaryTree(nums[index+1:])
	}
	return root
}

func maxOfNums(nums []int) (index, max int) {
	for i, num := range nums {
		if num > max {
			index, max = i, num
		}
	}
	return
}
