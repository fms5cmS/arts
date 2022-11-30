package greedy

import (
	"github.com/go-playground/assert/v2"
	"sort"
	"testing"
)

func getMinCoinCount(coins []int, k int) int {
	total, count := k, 0
	length := len(coins)
	sort.Ints(coins)
	for i := length - 1; i >= 0; i-- {
		curCount := total / coins[i] // 计算当前面值最多可以放几个
		total -= curCount * coins[i] // 计算余额
		count += curCount
		if total == 0 {
			return count
		}
	}
	return -1 // 未凑出
}

func TestGetMinCoinCount(t *testing.T) {
	tests := []struct {
		name  string
		coins []int
		k     int
		want  int
	}{
		{
			name:  "1",
			coins: []int{5, 3},
			k:     11,
			want:  3,
		},
		{
			name:  "2",
			coins: []int{1, 2, 5},
			k:     12,
			want:  3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, getMinCoinCount(test.coins, test.k))
		})
	}

}
