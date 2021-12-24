package binarySearch

// 左闭右闭区间，target 在 [left, right] 间
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 从 right 的初始值可以看出，left == right 是有意义的
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] > target {
			// 为了不重复判断，这里 right = middle-1，而不能使用 right = middle
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
	// left == right 是无意义的，会越界
	for left < right {
		middle := left + (right-left)>>1
		if nums[middle] > target {
			// 不能使用 right = middle-1，因为搜索时不会取值到右边界点，所以 middle-1 索引处的值永远不会被判断
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
