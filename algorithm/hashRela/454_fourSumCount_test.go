package hashRela

// 注意与 15 三数之和、18 四数之后 题的区别，15、18 不适合用哈希法，因为要求返回结果中不能包含重复的元素，会很复杂
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// key：nums1、nums2 中数值之和；value：nums1、nums2 中该数值之和出现的次数
	twoSumMap := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			twoSumMap[num1+num2]++
		}
	}
	// 统计 num1+num2+num3+num4=0 的次数
	count := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			count += twoSumMap[0-(num3+num4)]
		}
	}
	return count
}
