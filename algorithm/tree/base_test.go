package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestLevelPrint(t *testing.T) {
	array := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := constructTreeByArray(array)
	levels := root.LevelPrint()
	assert.Equal(t, array, levels)

	//fmt.Println(levels)
	//if len(levels) != len(array) {
	//	t.Fatal("length is not equal")
	//}
	//for i, v := range array {
	//	if v != levels[i] {
	//		t.Fatalf("not equal: %d, array(%v) VS levels(%v)", i, v, levels[i])
	//	}
	//}
}
