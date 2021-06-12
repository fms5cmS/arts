package treeRela

import (
	"fmt"
	"testing"
)

func TestGenGenerate(t *testing.T) {
	trees := generateTrees(3)
	for _, tree := range trees {
		fmt.Printf("%+v", tree)
	}
}