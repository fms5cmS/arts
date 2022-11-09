package binarySearch

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left, right := 1, num>>1
	for left <= right {
		mid := left + (right-left)>>1
		get := mid * mid
		if get == num {
			return true
		} else if get > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
