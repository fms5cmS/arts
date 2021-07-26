package arrayRela

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	k := len(nums) - 1
	// 双指针法，结果数组 result 的最大值在原数组 nums 的两端，要么最左，要么最右
	// 注：i <= j ！因为最后要处理两个元素
	for i, j := 0, len(nums); i <= j; {
		if nums[i]*nums[i] < nums[j]*nums[j] {
			result[k] = nums[j] * nums[j]
			j--
		} else {
			result[k] = nums[i] * nums[i]

			i++
		}
		k--
	}
	return result
}
