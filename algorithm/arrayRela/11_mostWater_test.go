package arrayRela

// 11. Container with most water
// 面积 = 底 * 高
// 其中双指针每次向内移动，所以底一定是减小的
// 要想得到最大的面积，必须高增大，所以每次最小的高向内移动
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0
	for l < r {
		if height[l] < height[r] {
			maxArea = maxInt((r-l)*height[l], maxArea)
			l++
		} else {
			maxArea = maxInt((r-l)*height[r], maxArea)
			r--
		}
	}
	return maxArea
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
