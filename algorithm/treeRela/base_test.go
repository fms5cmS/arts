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

func TestConstructTreeByLevelOrder(t *testing.T) {
	levelVals := []interface{}{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8}
	root := ConstructTreeByLevelOrder(levelVals)
	ret := LevelOrder(root)
	for _, ints := range ret {
		fmt.Println(ints)
	}
}
