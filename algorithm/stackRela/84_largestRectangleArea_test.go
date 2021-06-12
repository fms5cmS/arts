package stackRela

import "testing"

// 暴力解法：
// 遍历每一个高度，然后从当前高度分别向左右两侧扩散，寻找比当前高度高的，计算面积
// 时间复杂度 O(N^2)，空间复杂度 O(1)
func largestRectangleArea_force(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}
	ret := 0
	for i := 0; i < length; i++ {
		curHeight := heights[i]
		left := i
		for left > 0 && heights[left-1] >= curHeight {
			left--
		}
		right := i
		for right < length-1 && heights[right+1] >= curHeight {
			right++
		}
		width := right - left + 1
		ret = max(ret, curHeight*width)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func largestRectangleArea(heights []int) int {
// 	length := len(heights)
// 	if length == 0 {
// 		return 0
// 	}
// 	if length == 1 {
// 		return heights[0]
// 	}
// 	ret := 0
// 	stack := make([]int, 0)
// 	for i := 0; i < length; i++ {
// 		for len(stack) != 0 && heights[i] < heights[stack[len(stack)-1]] {
// 			curHeight := heights[stack[len(stack)-1]]
//
// 		}
// 	}
// }

func TestLargest(t *testing.T) {
	nums := []int{2, 1, 5, 6, 2, 3}
	ret := largestRectangleArea_force(nums)
	t.Log(ret)
}
