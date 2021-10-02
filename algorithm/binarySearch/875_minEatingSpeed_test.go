package binarySearch

import (
	"fmt"
	"testing"
)

func minEatingSpeed(piles []int, h int) int {
	// slow, fast, mid 都是吃香蕉的速度,最慢为 1
	slow, fast := 1, maxOfInts(piles)
	for slow <= fast {
		mid := slow + (fast-slow)>>1
		if totalTimeForEating(piles, mid) <= h {
			fast = mid - 1
		} else {
			slow = mid + 1
		}
	}
	return slow
}

func maxOfInts(numbers []int) int {
	max := 0
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

// 计算以 k 为速度，吃完 piles 需要的时间
func totalTimeForEating(piles []int, k int) int {
	time := 0
	for _, v := range piles {
		time += v / k
		if v%k > 0 {
			time++
		}
	}
	return time
}

func TestMinEatingSpeed(t *testing.T) {
	fmt.Println(minEatingSpeed([]int{312884470}, 968709470))
}
