package dpRela

// 1. 奇数：奇数的二进制表示一定比前一个偶数的二进制表示多 1（低位的 1）
// 2. 偶数：偶数的二进制表示一定和该数字右移一位后的二进制表示相等
func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if i%2 == 1 {
			result[i] = result[i-1] + 1
		} else {
			result[i] = result[i>>1]
		}
	}
	return result
}

func countBitsDP(n int) []int {
	result := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}
