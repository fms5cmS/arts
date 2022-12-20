package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 从 1～9 中找出 k 个数字，这些数字之和为 n。每个数字只能使用一次
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0, k)
	var backtracking func(sum, startValue int)
	backtracking = func(sum, startValue int) {
		// 剪枝操作
		if sum > n {
			return
		}
		if sum == n && len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := startValue; i < 9; i++ {
			sum += i
			path = append(path, i)
			backtracking(sum, i+1)
			// 回溯
			sum -= i
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 1)
	return result
}

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name string
		k    int
		n    int
		want [][]int
	}{
		{
			name: "1",
			k:    3,
			n:    7,
			want: [][]int{{1, 2, 4}},
		},
		{
			name: "2",
			k:    3,
			n:    9,
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name: "3",
			k:    4,
			n:    1,
			want: [][]int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, combinationSum3(test.k, test.n))
		})
	}
}
