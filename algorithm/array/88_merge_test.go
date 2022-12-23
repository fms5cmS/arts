package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 	从后往前比较！！
	i, j := m-1, n-1 // 分别为 nums1、nums2 真实数据的最后一个索引
	k := m + n - 1   // nums1 切片最后一个索引
	for ; j >= 0; k-- {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			name:  "1",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "2",
			nums1: []int{4, 5, 6, 0, 0, 0},
			m:     3,
			nums2: []int{1, 2, 3},
			n:     3,
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "3",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			name:  "4",
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			merge(test.nums1, test.m, test.nums2, test.n)
			assert.Equal(t, test.want, test.nums1)
		})
	}
}
