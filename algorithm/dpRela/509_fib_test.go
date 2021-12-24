package dpRela

import (
	"testing"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	// 1. dp 数组（下标对应的元素值对应了该下标的斐波那契数），因为存在 fib(0)，所以数组长度为 n+1
	dp := make([]int, n+1)
	// 3. dp 数组初始化，fib(0) = 0，fib(1) = 1
	dp[0], dp[1] = 0, 1
	// 2. 递推公式：fib(n) = fib(n-1) + fib(n-2)
	// 4. 遍历顺序：从前往后
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	// fmt.Println(dp) 打印示例的 dp 数组，判断是否符合预期
	return dp[n]
}

// 简化
func fibSimple(n int) int {
	if n <= 1 {
		return n
	}
	// 从下面递推公式可以看出，dp 数组只需要两个元素即可
	a, b := 0, 1
	sum := 0
	// 递推公式：fib(n) = fib(n-1) + fib(n-2)
	// 遍历顺序：从前往后
	for i := 2; i <= n; i++ {
		sum = a + b
		a, b = b, sum
	}
	return sum
}

func TestFib(t *testing.T) {
	t.Log(fib(5))
	t.Log(fibSimple(5))
}
