package binarySearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// left, right = 1, 8. middle = 4  16 > 8
// left, right = 1, 3. middle = 2  4 < 8
// left, right = 2, 3. middle = 2  4 < 8
// left, right = 3, 3. middle = 3  9 > 8
// left, right = 3, 2
func mySqrt(x int) int {
	left, right := 1, x
	for left <= right {
		mid := left + (right-left)>>1
		get := mid * mid
		if get == x {
			return mid
		} else if get > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 2, mySqrt(8))
}
