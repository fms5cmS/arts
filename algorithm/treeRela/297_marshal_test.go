package treeRela

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type Codec struct {
	// 节点的分隔符
	separatorStr string
	// 空指针时填充的字符
	nullStr string
}

// 原函数名为 Constructor()
func ConstructorForSerialize() Codec {
	return Codec{separatorStr: ",", nullStr: "@"}
}

// 采用前序遍历的方式来序列化
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return this.nullStr
	}
	return fmt.Sprintf("%d%s%s%s%s", root.Val, this.separatorStr, this.serialize(root.Left), this.separatorStr, this.serialize(root.Right))
}

// 一般情况下，单独的前/中/后序遍历结果是无法还原二叉树的，因为缺少空指针的信息，至少要得到前、中、后的两种遍历结果才能还原
// 而这里是有空指针信息的，所以可以还原
// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, this.separatorStr)
	return this.buildTreeByDeserialize(&nodes)
}

// 注意：这里的 strs 必须是切片的指针，以保证代码中 root.Left 取出一个字符处理后，后续 root.Right 不会重复处理
func (this *Codec) buildTreeByDeserialize(strs *[]string) *TreeNode {
	val := (*strs)[0]
	*strs = (*strs)[1:]
	if val == this.nullStr {
		return nil
	}
	rootVal, _ := strconv.Atoi(val)
	root := &TreeNode{Val: rootVal}
	root.Left = this.buildTreeByDeserialize(strs)
	root.Right = this.buildTreeByDeserialize(strs)
	return root
}

// 上面使用了前序遍历的方式来序列化和反序列化
// 也可以使用中、后、层序遍历的方式
// 记得补充层序遍历的方式
// 5,3,2,@,@,4,@,@,6,@,7,@,@
func TestSerialize(t *testing.T) {
	bst := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6,
			Right: &TreeNode{Val: 7}},
	}
	codec := ConstructorForSerialize()
	str := codec.serialize(bst)
	t.Log(str)
	t.Log(LevelOrder(codec.deserialize(str)))
}
