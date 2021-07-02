package arrayRela

//「使……最大值尽可能小」是二分搜索题目常见的问法。
// 使用二分法：
// 如果 m == 1，则整个 nums 就是一个子数组，返回 nums 的和
// 如果 m == len(nums)，则每个 nums 的元素作为一个子数组，这 m 个子数组各自和的最大值就是 nums 的元素最大值
// 对于其他有效的 m 值，返回的值必定在上面的范围之内
// 示例：
// nums = [1, 2, 3, 4, 5], m = 3
// left = 5, right = 15  ===>  mid = 10，所以需要找出和最大且小于等于 10 的子数组的个数
// [1, 2, 3, 4], [5]  无法分为 3(m) 组，说明 mid 偏大 ===>  right = mid，继续查找
// mid = 7，找出和最大且小于等于7的子数组的个数，[1,2,3], [4], [5]，成功的找出了三组，说明 mid 可以继续降低 ===> right = mid
// mid = 6 ...
// mid = 5，再次找出和最大且小于等于5的子数组的个数，[1,2], [3], [4], [5]，发现有4组，此时的 mid 太小了，应该增大 mid，让 left=mid+1，此时 left=6，right=6，循环退出了
func splitArray(nums []int, m int) int {
	// 计算二分法左右两个端点，分别为 nums 的最大值、nums 的和
	left, right := 0, 0
	for _, num := range nums {
		right += num
		if num > left {
			left = num
		}
	}
	for left < right {
		mid := left + (right-left)>>1
		if canSplit(nums, mid, m) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// sumOfSubArr 表示当前分割子数组的和，totalSubArr 表示已经分割出的子数组的数量（包括当前子数组）
func canSplit(nums []int, x, m int) bool {
	sumOfSubArr, totalSubArr := 0, 1
	for _, num := range nums {
		if sumOfSubArr+num > x {
			totalSubArr++
			sumOfSubArr = num
		} else {
			sumOfSubArr += num
		}
	}
	return totalSubArr <= m
}
