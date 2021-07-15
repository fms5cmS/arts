package dpRela

// 找是否可以将这个数组分割成两个子集，使得两个子集的元素和相等
// 所以，只要找到集合里能够出现 sum / 2 的子集总和，就算是可以分割成两个相同元素和子集了。
// 类比 01 背包问题：背包容量 sum/2，要放入的商品为 nums 元素的值，背包中每个元素不可重复放入
func canPartition(nums []int) bool {
	sum := 0
	// 1. dp 数组，dp[i] 代表容量为 i 的背包，最大可以凑成 i 的子集之和为dp[i]
	// 题目中：1 <= nums.length <= 200 且 1 <= nums[i] <= 100
	// 总和不会大于 20000，背包最大容量为其一半 10001 即可
	// 3. 初始化，这里将所有下标元素均初始化为 0
	dp := make([]int, 10001)
	// 计算背包最大容量
	for _, num := range nums {
		sum += num
	}
	// 如果物品总重为奇数，则必然不可分割
	if sum%2 == 1 {
		return false
	}
	target := sum >> 1
	// 4. 遍历顺序，见 https://mp.weixin.qq.com/s/M4uHxNVKRKm5HPjkNZBnFA
	for _, num := range nums {
		for j := target; j >= num; j-- {
			// 2. 递推公式
			dp[j] = maxOf2Ints(dp[j], dp[j-num]+num)
		}
	}
	if dp[target] == target {
		return true
	}
	return false
}
