package sort

// 选择排序：每次从无序区找到最小元素放到有序区末尾
func selectionSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		// 查找 [i, len-1] 范围内最小值的索引
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		// 每次移动i，最多仅发生一次交换操作
		nums[minIndex], nums[i] = nums[i], nums[minIndex]
	}
}
