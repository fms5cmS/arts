package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 后序遍历 postorder 最后一个值为根节点
// 以上面根节点的值在中序遍历 inorder 中查找元素，该元素左侧为根节点左子树的元素，右侧为根节点右子树的元素
// 分割中序数组：找到根节点在中序数组的位置，左侧为左子树的元素，右侧为右子树的元素
// 分割后序数组：后序数组不像中序数组那样有明确的分割点，不过分割后序数组时一个重要点：
//     中序数组的大小一定要和后序数组的大小相同！！！
//   所以，后序数组可以按照中序数组的大小分割
func buildTreeByInAndPostOrder(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	// 后序数组最右侧元素为根节点的值
	length := len(postorder)
	root := &TreeNode{Val: postorder[length-1]}
	// 找到根节点值在中序数组中的位置
	index := -1
	for i, v := range inorder {
		if v == root.Val {
			index = i
		}
	}
	if index >= 0 {
		// 中序数组有明确的分割点：根节点将中序数组分割为左右子树
		// 注：后序数组没有明确的分割点，但左右子树在后续数组的长度一定与在中序数组中的长度相等
		root.Left = buildTreeByInAndPostOrder(inorder[:index], postorder[:index])
		root.Right = buildTreeByInAndPostOrder(inorder[index+1:], postorder[index:length-1])
		// 与 105 不同，这里要用右子树的长度来划分
		// rightLength := len(inorder[index+1:])
		// root.Left = buildTreeByInAndPostOrder(inorder[0:index], postorder[0:length-rightLength-1])
		// root.Right = buildTreeByInAndPostOrder(inorder[index+1:], postorder[length-rightLength-1:length-1])
	}
	return root
}

func TestBuildTreeByInAndPostOrder(t *testing.T) {
	tests := []struct {
		name      string
		inorder   []int
		postorder []int
		wantLevel []interface{}
	}{
		{
			name:      "first",
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			wantLevel: []interface{}{3, 9, 20, nil, nil, 15, 7},
		},
		{
			name:      "second",
			inorder:   []int{-1},
			postorder: []int{-1},
			wantLevel: []interface{}{-1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := buildTreeByInAndPostOrder(test.inorder, test.postorder)
			assert.Equal(t, root.LevelPrint(), test.wantLevel)
		})
	}
}
