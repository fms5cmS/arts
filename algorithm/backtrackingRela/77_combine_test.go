package backtrackingRela

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0, k) // 暂存符合条件的结果
	// startValue 是用来标记前面的值已经被用过了不能再用
	var backtracking func(startValue int)
	backtracking = func(startValue int) {
		// path 长度满足时，递归结束
		if len(path) == k {
			// 这里需要重新定义一个新的变量来接收符合条件的结果，直接将 path 放入 result 的话，由于切片底层数组的值一直在变化，会导致结果错误
			// 注意，这里 temp 的长度为 k，是因为 copy 函数会选择源切片、目标切片最短的长度来复制值
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}
		// for 循环来选择当前处理的值
		// for i := startValue; i <= n-(k-len(path))+1; i++ { // 剪枝操作，n=4，k=4 时打印 path 理解
		for i := startValue; i <= n; i++ {
			path = append(path, i)
			fmt.Println(path)
			// 递归函数用来从未被使用过的值（就需要用到 startValue 了）中找下一个值
			backtracking(i + 1)
			path = path[:len(path)-1] // 回溯，撤回上一步处理的值
		}
	}
	backtracking(1)
	return result
}

func TestCombine77(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want [][]int
	}{
		{
			name: "1",
			n:    4,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "2",
			n:    1,
			k:    1,
			want: [][]int{{1}},
		},
		{
			name: "3",
			n:    4,
			k:    4,
			want: [][]int{{1, 2, 3, 4}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, combine(test.n, test.k))
		})
	}
}
