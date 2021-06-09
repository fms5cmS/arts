package treeRela

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	a := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	ret := LevelOrder(a)
	for _, ints := range ret {
		fmt.Println(ints)
	}
}