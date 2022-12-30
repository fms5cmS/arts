package stackAndQueue

import (
	"arts/algorithm/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 暴力解法：
// 遍历每一个高度，然后从当前高度分别向左右两侧扩散，找到以该高度为最低高度的面积
// 时间复杂度 O(N^2)，空间复杂度 O(1)
func largestRectangleArea_force(heights []int) int {
	maxArea := 0
	for i, height := range heights {
		// 向左侧扩散
		left := i
		for left > 0 && heights[left-1] >= height {
			left--
		}
		// 向右侧扩散
		right := i
		for right < len(heights)-1 && heights[right+1] >= height {
			right++
		}
		// 计算最大面积
		width := right - left + 1
		maxArea = utils.Max(maxArea, height*width)
	}
	return maxArea
}

// https://leetcode.com/problems/largest-rectangle-in-histogram/solutions/28900/short-and-clean-o-n-stack-based-java-solution/
// 其实是对上面暴力解法的优化，通过一个栈来避免每个高度向左侧扩散时重复计算的逻辑
// 借助了一个栈，栈中存储了 hegiths 的索引，且索引对应的值是递增的！！
// 当要想栈中放入的索因所对应的值 height 大于栈顶元素时，就说明栈中那些比 height 大的值的右边界也都确定了，就可以将其弹出计算面积
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0, len(heights))
	maxArea := 0
	// 注意，这里的范围 i 最大是 len(heights)，是为了下面填充一个理论上的最小高度 0，从而将栈中所有的元素对应的高度都计算一遍
	for i := 0; i <= len(heights); i++ {
		// 初始化一个理论最低高度 0
		curHeight := 0
		if i != len(heights) {
			curHeight = heights[i]
		}
		// 当前遍历的高度大于等于栈顶元素对应的高度，则入栈
		if len(stack) == 0 || curHeight >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			// 说明栈顶元素对应的 height 的右侧边界已经找到了，计算其面积并与之前的 maxArea 比较后更新maxArea
			height := heights[stack[len(stack)-1]]
			// 栈顶元素弹出，因为目前已经在计算栈顶元素对应高度的最大面积了
			stack = stack[:len(stack)-1]
			// 计算宽度
			width := i
			if len(stack) != 0 {
				width = i - 1 - stack[len(stack)-1]
			}
			maxArea = utils.Max(maxArea, height*width)
			i--
		}
	}
	return maxArea
}

func TestLargest(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "1",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		{
			name:    "2",
			heights: []int{3, 5, 1, 7, 5, 9},
			want:    15,
		},
		{
			name:    "3",
			heights: []int{6, 7, 5, 2, 4, 5, 9, 3},
			want:    16,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, largestRectangleArea_force(test.heights))
			assert.Equal(t, test.want, largestRectangleArea(test.heights))
		})
	}
}
