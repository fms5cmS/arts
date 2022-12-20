package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 对称二叉树，必须是镜像对称
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compareIsSymmetric(root.Left, root.Right)
}

func compareIsSymmetric(left, right *TreeNode) bool {
	// 处理简单情况
	if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right == nil {
		return true
		//	值的比较必须放在最后，否则可能会空指针 panic，上面三种情况已经排除了 left、right 存在 nil 的情况
	} else if left.Val != right.Val {
		return false
	}
	// 走到这里时，left、right 均不为空，且值相等，需要比较子节点，注意下面比较的节点！
	outside := compareIsSymmetric(left.Left, right.Right)
	inside := compareIsSymmetric(left.Right, right.Left)
	return outside && inside
}

// 非递归
func isSymmetricByQueue(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, []*TreeNode{root.Left, root.Right}...)
	// 把左右两个紫薯要比较的元素顺序放进容器中，然后成对比较
	for len(queue) > 0 {
		leftNode, rightNode := queue[0], queue[1]
		queue = queue[2:]
		if leftNode == nil && rightNode == nil {
			continue
		}
		// 此时 leftNode、rightNode 至少有一个不为空
		// 如果左右有一个节点为空，或都不为空时两个节点值不等，返回 false
		if leftNode == nil || rightNode == nil || (leftNode.Val != rightNode.Val) {
			return false
		}
		queue = append(queue, []*TreeNode{leftNode.Left, rightNode.Right, leftNode.Right, rightNode.Left}...)
	}
	return true
}

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name  string
		array []interface{}
		want  bool
	}{
		{
			name:  "first",
			array: []interface{}{1, 2, 2, 3, 4, 4, 3},
			want:  true,
		},
		{
			name:  "second",
			array: []interface{}{1, 2, 2, nil, 3, nil, 3},
			want:  false,
		},
		{
			name:  "third",
			array: []interface{}{1, 2, 3},
			want:  false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := constructTreeByArray(test.array)
			assert.Equal(t, test.want, isSymmetric(root))
			assert.Equal(t, test.want, isSymmetricByQueue(root))
		})
	}
}
