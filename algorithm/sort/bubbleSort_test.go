package sort

// 冒泡排序：相邻的元素两两比较，根据大小来交换元素的位置
func bubbleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	length := len(nums)
	for i := length - 1; i >= 1; i-- {
		sorted := true // 假定未排序区间内元素有序
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				sorted = false // 发生了交换，则说明未排序区间内元素无序
			}
		}
		if sorted {
			break // 未排序区间元素有序，所有元素已排序完成，退出循环
		}
	}
}
