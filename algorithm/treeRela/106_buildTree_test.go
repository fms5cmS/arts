package treeRela

import (
	"fmt"
	"testing"
)

// 后序遍历 postorder 最后一个值为根节点
// 以上面根节点的值在中序遍历 inorder 中查找元素，该元素左侧为根节点左子树的元素，右侧为根节点右子树的元素
func buildTreeByInAndPostOrder(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	length := len(postorder)
	root := new(TreeNode)
	// 对根节点赋值
	root.Val = postorder[length-1]
	// 找到根节点的值在 inorder 中的索引，从而确定 inorder 中左右子树各自的元素
	index := getIndex(inorder, postorder[length-1])
	if index >= 0 {
		// 与 105 不同，这里要用右子树的长度来划分
		rightLength := len(inorder[index+1:])
		root.Left = buildTreeByInAndPostOrder(inorder[0:index], postorder[0:length-rightLength-1])
		root.Right = buildTreeByInAndPostOrder(inorder[index+1:], postorder[length-rightLength-1:length-1])
	}
	return root
}

func TestBuildTreeByInAndPostOrder(t *testing.T) {
	//postorder := []int{9, 15, 7, 20, 3}
	//inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{2, 1}
	inorder := []int{1, 2}
	tree := buildTreeByInAndPostOrder(inorder, postorder)
	treeArr := LevelOrder(tree)
	for _, level := range treeArr {
		fmt.Println(level)
	}
}
