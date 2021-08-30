package hashRela

import (
	"fmt"
	"testing"
)

func isHappy(n int) bool {
	set := make(map[int]bool)
	// n == 1 代表是快乐数
	// 如果这个值出现过（set[n] == true）,说明进入了无限循环
	for n != 1 && !set[n] {
		n, set[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func TestIsHappy(t *testing.T) {
	fmt.Println(isHappy(19))
}
