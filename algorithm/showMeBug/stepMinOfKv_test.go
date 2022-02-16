package showMeBug

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// StepMinOfKv 找到 key 和 value 最接近的 key
// 如果有多个，返回最小的 key
func StepMinOfKv(src map[int]int) int {
	result := math.MaxInt64
	minSub := math.MaxInt64
	for key, value := range src {
		sub := abs(key - value)
		if sub < minSub {
			result = key
			minSub = sub
		} else if sub == minSub && key < result {
			result = key
		}
	}
	return result
}

func abs(src int) int {
	if src < 0 {
		src = -src
	}
	return src
}

func TestStepMinOgKv(t *testing.T) {
	src := map[int]int{
		1:  93,
		10: 55,
		15: 30,
		20: 19,
		23: 11,
		30: 2,
	}
	assert.Equal(t, 20, StepMinOfKv(src))
}
