package arrayRela

// 283. Move zeros
// boundary、explore 两个指针同时从左侧移动。boundary 作为 0 与非 0 元素的分界线
// 当两者遇到 0 时，开始分离，boundary 不再移动，而 explore 继续移动
// explore 遇到 0 时，与 boundary(此时指向 0) 交换元素，然后两个指针同时移动
func moveZeroes(nums []int) {
	boundary, explore := 0, 0
	for explore < len(nums) {
		if nums[explore] != 0 {
			nums[boundary], nums[explore] = nums[explore], nums[boundary]
			boundary++
			explore++
		} else {
			explore++
		}
	}
}
