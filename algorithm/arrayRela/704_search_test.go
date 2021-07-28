package arrayRela

// 左闭右闭区间，target 在 [left, right] 间
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

// 左闭右开区间，target 在 [left, right) 间
func search2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		middle := left + (right-left)>>1
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
