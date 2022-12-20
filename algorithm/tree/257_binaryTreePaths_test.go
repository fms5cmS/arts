package tree

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

// 因为是从根节点到叶子节点的路径，需要父节点指向子节点，所以必须是**前序遍历**
// 实现还需要回溯来回退一个路径再进入另一个路径
func binaryTreePaths(root *TreeNode) []string {
	path, result := make([]string, 0), make([]string, 0)
	if root == nil {
		return result
	}
	traversal(root, &path, &result)
	return result
}

// 递归 + 回溯
func traversal(cur *TreeNode, path *[]string, result *[]string) {
	*path = append(*path, strconv.Itoa(cur.Val))
	if cur.Left == nil && cur.Right == nil {
		*result = append(*result, strings.Join(*path, "->"))
	}
	if cur.Left != nil {
		traversal(cur.Left, path, result)
		*path = (*path)[:len(*path)-1] // 回溯
	}
	if cur.Right != nil {
		traversal(cur.Right, path, result)
		*path = (*path)[:len(*path)-1] // 回溯
	}
}

func TestBinaryTreePaths(t *testing.T) {
	tests := []struct {
		name      string
		leveOrder []interface{}
		want      []string
	}{
		{
			name:      "first",
			leveOrder: []interface{}{1, 2, 3, nil, 5},
			want:      []string{"1->2->5", "1->3"},
		},
		{
			name:      "second",
			leveOrder: []interface{}{1},
			want:      []string{"1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.leveOrder)
			assert.Equal(t, test.want, binaryTreePaths(root))
		})
	}
}
