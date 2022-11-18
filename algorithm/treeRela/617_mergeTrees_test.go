package treeRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 如何同时遍历两个二叉树呢？和遍历一个树逻辑是一样的，只不过传入两个树的节点，同时操作。
// 递归，前序遍历（中、后序也可以）
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root := new(TreeNode)
	root.Val = root1.Val + root2.Val
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

// 迭代
// 思路同 101，把两棵树的节点都入队再处理
func mergeTreesNotRecursion(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue := []*TreeNode{root1, root2}
	for len(queue) > 0 {
		node1, node2 := queue[0], queue[1]
		queue = queue[2:]
		// 两个节点一定不为空，val 相加
		node1.Val += node2.Val
		// 两棵树的左右节点不为空，入队
		if node1.Left != nil && node2.Left != nil {
			queue = append(queue, node1.Left, node2.Left)
		}
		if node1.Right != nil && node2.Right != nil {
			queue = append(queue, node1.Right, node2.Right)
		}
		// 由于是将 root2 merge 到了 root1 上，所以仅需判断 root1 比 root2 缺少的部分，并将 root2 对应的部分直接挂到 root1 上即可
		// 注意：这里没有入队操作
		// node1 的左节点为空，node2 的不为空，将 node2 的值赋值过去
		if node1.Left == nil && node2.Left != nil {
			node1.Left = node2.Left
		}
		// node1 的右节点为空，node2 的不为空，将 node2 的值赋值过去
		if node1.Right == nil && node2.Right != nil {
			node1.Right = node2.Right
		}
	}
	return root1
}

func TestMergeTrees(t *testing.T) {
	tests := []struct {
		name  string
		root1 []interface{}
		root2 []interface{}
		want  []interface{}
	}{
		{
			name:  "first",
			root1: []interface{}{1, 3, 2, 5},
			root2: []interface{}{2, 1, 3, nil, 4, nil, 7},
			want:  []interface{}{3, 4, 5, 5, 4, nil, 7},
		},
		{
			name:  "second",
			root1: []interface{}{1},
			root2: []interface{}{1, 2},
			want:  []interface{}{2, 2},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root1 := constructTreeByArray(test.root1)
			root2 := constructTreeByArray(test.root2)
			get1 := mergeTrees(root1, root2)
			get2 := mergeTreesNotRecursion(root1, root2)

			assert.Equal(t, test.want, get1.LevelPrint())
			assert.Equal(t, test.want, get2.LevelPrint())
		})
	}
}
