package array

import (
	"arts/algorithm/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 11. Container with most water
// 面积 = 底 * 高
// 其中双指针每次向内移动，所以底一定是减小的
// 要想得到最大的面积，必须高增大，所以每次最小的高向内移动
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		if height[left] < height[right] {
			maxArea = utils.Max((right-left)*height[left], maxArea)
			left++
		} else {
			maxArea = utils.Max((right-left)*height[right], maxArea)
			right--
		}
	}
	return maxArea
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "3",
			height: []int{1, 10, 6, 2, 5, 4, 10, 3, 7},
			want:   50,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, maxArea(test.height))
		})
	}
}
