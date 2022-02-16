package showMeBug

import (
	"fmt"
	"math"
	"testing"
)

// MaxMultiply 找到给定一个整数数组中具有最大乘积的相邻元素对并返回该乘积
func MaxMultiply(src []int) int {
	if len(src) == 1 {
		return src[0]
	}
	first, second := 0, 1
	result := math.MinInt64
	for ; second < len(src); first, second = first+1, second+1 {
		if src[first]*src[second] > result {
			result = src[first] * src[second]
		}
	}
	return result
}

func TestMaxMultiply(t *testing.T) {
	fmt.Println(MaxMultiply([]int{3, 6, -2, -5, 7, 3}))
	fmt.Println(MaxMultiply([]int{-2, -1}))
}
