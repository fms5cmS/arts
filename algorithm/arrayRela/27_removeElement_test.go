package arrayRela

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if val != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
