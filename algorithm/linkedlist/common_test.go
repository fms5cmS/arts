package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	array1 := []int{1, 2, 6, 3, 4, 5, 6}
	list1 := generateListViaArray(array1)
	t.Logf("list1: %s", list1)
	arr1 := generateArrayViaList(list1)
	assert.Equal(t, array1, arr1)

	array2 := []int{}
	list2 := generateListViaArray(array2)
	t.Logf("list2: %s", list2)
	arr2 := generateArrayViaList(list2)
	assert.Equal(t, array2, arr2)

	array3 := []int{1}
	list3 := generateListViaArray(array3)
	t.Logf("list3: %s", list3)
	arr3 := generateArrayViaList(list3)
	assert.Equal(t, array3, arr3)
}
