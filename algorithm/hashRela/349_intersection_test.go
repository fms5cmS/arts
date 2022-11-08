package hashRela

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	result := make([]int, 0)
	for _, num := range nums1 {
		set[num] = struct{}{}
	}
	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			result = append(result, num)
			delete(set, num)
		}
	}
	return result
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		result []int
	}{
		{
			nums1:  []int{1, 2, 2, 1},
			nums2:  []int{2, 2},
			result: []int{2},
		},
		{
			nums1:  []int{4, 9, 5},
			nums2:  []int{9, 4, 9, 8, 4},
			result: []int{4, 9},
		},
	}
	for _, test := range tests {
		result := intersection(test.nums1, test.nums2)
		sort.Ints(result)
		assert.Equal(t, test.result, result)
	}
}
