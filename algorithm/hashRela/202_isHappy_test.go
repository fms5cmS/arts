package hashRela

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1. 如何取数值各个位上的数
// 2. 出现无限循环就代表一定不是 happy number
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

func TestGetSum(t *testing.T) {
	tests := []struct {
		n   int
		sum int
	}{
		{
			n:   19,
			sum: 82,
		},
		{
			n:   82,
			sum: 68,
		},
		{
			n:   68,
			sum: 100,
		},

		{
			n:   100,
			sum: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.sum, getSum(test.n))
	}
}

func TestIsHappy(t *testing.T) {
	tests := []struct {
		n     int
		happy bool
	}{
		{
			n:     19,
			happy: true,
		},
		{
			n:     2,
			happy: false,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.happy, isHappy(test.n))
	}
}
