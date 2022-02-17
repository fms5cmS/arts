package sort

// 快速排序
// 1. 对数组**分区**，小于原数组最后一个值的放在左侧，大于的放在右侧；
// 2. 对左侧和右侧分别继续进行经典快排
func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	sortOfQuickSort(nums, 0, len(nums))
}

func sortOfQuickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	q := partitionOfQuickSort(nums, low, high)
	sortOfQuickSort(nums, low, q-1)
	sortOfQuickSort(nums, q+1, high)
}

// 获取分区点的索引
func partitionOfQuickSort(nums []int, low, high int) int {
	partitionPoint := nums[high] // 以最后一个元素作为分区点
	i := low
	// j 从左往右开始移动，当 nums[j] < pivot 时，交换索引 i 和 j 各自对应的值
	for j := low; j < high; j++ {
		if nums[j] < nums[partitionPoint] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	// 此时索引 i 对应的值就是从左往右第一个大于 partitionPoint 的值
	// 交换 i 和 high(即 partitionPoint) 对应的值，则 i 就是分区点的索引
	nums[i], nums[high] = nums[high], nums[i]
	return i
}
