package arrayRela

import "testing"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	index, move := 0, 1
	for ; move < len(nums); move++ {
		if nums[move] != nums[index] {
			nums[index+1] = nums[move]
			index++
		}
	}
	return index+1
}

func removeDuplicates_slow(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	boundary, explore := 0, 1
	for ; explore < len(nums); explore++ {
		if nums[explore] != nums[boundary] {
			boundary++
			nums[explore], nums[boundary] = nums[boundary], nums[explore]
		}
	}
	return boundary + 1
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	i := removeDuplicates(nums)
	t.Log(i)
	t.Log(nums[0:i])
}
