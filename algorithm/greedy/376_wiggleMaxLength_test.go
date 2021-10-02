package greedy

// 贪心算法
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	curDiff, preDiff := 0, 0
	result := 1
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		// 注意：不能使用 curDiff*preDiff<=0 作为判断条件，因为必须保证 curDiff 不等于 0
		// 而里面更新 preDiff 又会保证只有边界条件最开始的时候会出现 preDiff = 0
		if (curDiff > 0 && preDiff <= 0) || (preDiff >= 0 && curDiff < 0) {
			result++
			preDiff = curDiff
		}
	}
	return result
}

func wiggleMaxLength2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	curDiff, preDiff := 0, 0
	result := 1
	i := 1
	// 找到第一个差值不为 0 的元素
	for ; i < len(nums); i++ {
		preDiff = nums[i] - nums[i-1]
		if preDiff != 0 {
			// 这里注意 result 需要加一
			result++
			break
		}
	}
	for ; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if curDiff*preDiff < 0 {
			result++
			preDiff = curDiff
		}
	}
	return result
}
