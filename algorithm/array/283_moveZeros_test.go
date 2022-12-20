package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}

func moveZeroes2(nums []int) {
	lastNonZeroIndex, cur := 0, 0
	for ; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[lastNonZeroIndex], nums[cur] = nums[cur], nums[lastNonZeroIndex]
			lastNonZeroIndex++
		}
	}
}

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "2",
			nums: []int{0},
			want: []int{0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			nums1, nums2 := make([]int, len(test.nums)), make([]int, len(test.nums))
			copy(nums1, test.nums)
			copy(nums2, test.nums)
			moveZeroes(nums1)
			assert.Equal(t, test.want, nums1)

			moveZeroes2(nums2)
			assert.Equal(t, test.want, nums2)
		})
	}
}
