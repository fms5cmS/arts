package arrayRela

import "testing"

// 70. Climbing Stairs
// 动态规划
// 					n,   0 < n < 3
// f(n) =   0,   n < 0
// 					f(n-1) + f(n-2),  n > 3
func climbStairs(n int) int {  // n 是给定的正整数
	if n <= 2 {
		return n
	}
	i, j := 1, 2
	num := 0
	for k := 3; k <= n; k++ {
		num = i + j
		i, j = j, num
	}
	return num
}

func TestClimbStairs(t *testing.T) {
	t.Log(climbStairs(3))
}
