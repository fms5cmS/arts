package binarySearch

import (
	"fmt"
	"testing"
)

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
	fmt.Println(mySqrt(4))
}