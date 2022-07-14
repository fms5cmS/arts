package slideWindow

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 暴力解法
func minSubArrayLenForce(target int, nums []int) int {
	result := math.MaxInt32
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				length := j - i + 1
				if length < result {
					result = length
					break
				}
			}
		}
	}
	if result == math.MaxInt32 {
		result = 0
	}
	return result
}

// O(n)
func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt32
	// 滑动窗口数值之和 sum，起始位置 start，长度 length
	sum, start, length := 0, 0, 0
	// 窗口范围为 [start, end]
	for end, num := range nums {
		sum += num
		// 当前窗口数值之和大于 target 时，窗口起始位置向前移动，注意这里是 for 循环！
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

func TestMinSubArray(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{
			nums:   []int{2, 3, 1, 2, 4, 3},
			target: 7,
			output: 2,
		},
		{
			nums:   []int{1, 4, 4},
			target: 4,
			output: 1,
		},
		{
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			target: 11,
			output: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, minSubArrayLenForce(test.target, test.nums))
		assert.Equal(t, test.output, minSubArrayLen(test.target, test.nums))
	}
}
