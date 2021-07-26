package arrayRela

import "math"

// O(n)
func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt32
	// 滑动窗口数值之和 sum，起始位置 start，长度 length
	sum, start, length := 0, 0, 0
	// 窗口范围为 [start, end]
	for end, num := range nums {
		sum += num
		// 当前窗口数值之和大于 target 时，窗口起始位置向前移动
		for sum >= target {
			length = end - start + 1
			if result > length {
				result = length
			}
			// start 向前移动，sum 的值需要将 start 当前的值减去
			sum -= nums[start]
			start++
		}
	}
	if result == math.MaxInt32 {
		result = 0
	}
	return result
}
