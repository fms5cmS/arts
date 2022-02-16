package showMeBug

import (
	"fmt"
	"testing"
)

// Tree2List 将拥有任意数目子节点的树转为列表。其中，树中节点的兄弟节点也要位于列表中数据的相邻位置。
//       1
//      / \
//     2   3  -> [1,2,3,4,5,6,7]
//    /|\   \
//   4 5 6   7
func Tree2List(root *Node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			result = append(result, cur.Value)
			queue = append(queue, cur.Children...)
		}
	}
	return result
}

func TestTree2List(t *testing.T) {
	root := &Node{Value: 1, Children: []*Node{
		{Value: 2, Children: []*Node{
			{Value: 4},
			{Value: 5},
			{Value: 6},
		}},
		{Value: 3, Children: []*Node{{Value: 7, Children: nil}}},
	}}
	fmt.Println(Tree2List(root))
}
