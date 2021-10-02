package arrayRela

import (
	"fmt"
	"sort"
	"testing"
)

// 每次排序后，仅针对数组的最小值进行取反
// 共计执行上面操作 k 次
func largestSumAfterKNegations(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < k; i++ {
		sort.Ints(nums)
		nums[0] = -nums[0]
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// 注意，取反操作不能放到 if else 的上面来统一处理！
func largestSumAfterKNegations2(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	sort.Ints(nums)
	// 如果数组的元素都是大于等于 0 的，那就仅针对第一个值做 k 次取反操作
	if nums[0] >= 0 {
		for i := 0; i < k; i++ {
			nums[0] = ^nums[0] + 1
		}
	} else {
		for i := 0; i < k; i++ {
			if nums[0] >= 0 {
				// 按位取反再加一就是这个数的相反数
				nums[0] = ^nums[0] + 1
				continue
			} else {
				nums[0] = ^nums[0] + 1
				sort.Ints(nums)
			}
		}
	}

	for _, num := range nums {
		sum += num
	}
	return sum
}

func TestLargestSumAfterKNegation(t *testing.T) {
	nums := []int{3, -1, 0, 2} // k = 3, sum =
	//nums := []int{4, 2, 3} // k = 1sum = 5
	//nums := []int{2, -3, -1, 5, -4} // 2
	fmt.Println(largestSumAfterKNegations2(nums, 3))
}
