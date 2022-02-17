package sort

// 归并排序
func mergeSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	partitionOfMergeSort(nums, 0, len(nums)-1)
}

func partitionOfMergeSort(nums []int, start, end int) {
	if start > end {
		return
	}
	middle := start + (end-start)>>1
	partitionOfMergeSort(nums, start, middle)
	partitionOfMergeSort(nums, middle+1, end)
	mergeOfMergeSort(nums, start, middle, end)
}

func mergeOfMergeSort(nums []int, start, middle, end int) {
	temp := make([]int, end-start+1)
	pointA, pointB := start, middle+1
	i := 0
	// 分别比较两个区间内的元素大小，按序放入辅助数组
	for pointA <= middle && pointB <= end {
		if nums[pointA] <= nums[pointB] {
			temp[i] = nums[pointA]
			pointA++
		} else {
			temp[i] = nums[pointB]
			pointB++
		}
		i++
	}
	// 其中一个区间的元素已遍历完，将另一个区间的元素放入辅助数组
	for pointA <= middle {
		temp[i] = nums[pointA]
		i++
		pointA++
	}
	for pointB <= end {
		temp[i] = nums[pointB]
		i++
		pointB++
	}
	// 辅助数组的数据放回原数组的指定区间内
	copy(nums[start:end+1], temp) // 由于是左闭右开区间，所以是 [start:end+1]
}
