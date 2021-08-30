package hashRela

func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	recordSet := make(map[int]int)
	// 遍历 nums1，初始值均为 1
	for _, num := range nums1 {
		recordSet[num] = 1
	}
	for _, num := range nums2 {
		// 通过 count 实现重复值只去一次放入 result
		if count, exists := recordSet[num]; exists && count > 0 {
			result = append(result, num)
			recordSet[num]--
		}
	}
	return result
}
