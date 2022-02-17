package sort

// 插入排序：想象玩扑克抽牌的过程，每抽到一张牌 x 就和前面的牌比较，
//   1. x 大则不做任何操作，并结束
//   2. x 小则交换，然后再和更前面的牌比较，再进行 1 或 2 的操作
func insertionSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	// [0,i)是有序区
	// 无序区的第一个元素依次与有序区的元素从后往前比较
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
