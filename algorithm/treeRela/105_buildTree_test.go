package treeRela

import (
	"fmt"
	"testing"
)

// 前序遍历 preorder 第一个值为根节点
// 以上面根节点的值在中序遍历 inorder 中查找元素，该元素左侧为根节点左子树的元素，右侧为根节点右子树的元素
// 分割中序数组：找到根节点在中序数组的位置，左侧为左子树的元素，右侧为右子树的元素
// 分割前序数组：前序数组不像中序数组那样有明确的分割点，不过分割前序数组时一个重要点：
//     中序数组的大小一定要和前序数组的大小相同！！！
//   所以，前序数组可以按照中序数组的大小分割
func buildTreeByPreAndInOrder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	// 对根节点赋值
	root.Val = preorder[0]
	// 找到根节点的值在 inorder 中的索引，从而确定 inorder 中左右子树各自的元素
	index := getIndex(inorder, preorder[0])
	if index >= 0 {
		root.Left = buildTreeByPreAndInOrder(preorder[1:index+1], inorder[:index])
		root.Right = buildTreeByPreAndInOrder(preorder[index+1:], inorder[index+1:])
		// 根据前序遍历的特点，根节点元素后面跟的是左子树的元素，所以这里 preorder 的范围要用左子树元素的数量 len(inorder[:index]) 来确定
		// 又因为区间是左闭右开区间，所以为 1:len(inorder[:index])+1，而 len(inorder[:index])+1 之后的均为右子树元素
		// root.Left = buildTreeByPreAndInOrder(preorder[1:len(inorder[:index])+1], inorder[:index])
		// root.Right = buildTreeByPreAndInOrder(preorder[len(inorder[:index])+1:], inorder[index+1:])
	}
	return root
}

func TestBuildTreeByPreAndInOrder(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTreeByPreAndInOrder(preorder, inorder)
	treeArr := LevelOrder(tree)
	for _, level := range treeArr {
		fmt.Println(level)
	}
}
