package hashRela

func twoSum(nums []int, target int) []int {
	recordMap := make(map[int]int)
	for i, num := range nums {
		if index, exists := recordMap[target-num]; exists {
			return []int{index, i}
		} else {
			recordMap[num] = i
		}
	}
	return []int{}
}
