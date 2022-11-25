package backtrackingRela

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// cpu的亲和度算法：
// cpu分为两组[0,1,2,3] [4,5,6,7]
// 请求分配cpu的参数为 1，2，4，8
//
// cpu的亲和度值：
// 请求1个cpu时，cpu的优先级为1，3，2，4。 注释：优先选择仅剩 1 个可用 cpu 的组，如果不存在的话，选择仅剩 3、2、4 个的组
// 请求2个cpu时，cpu的优先级为2，4，3
// 请求4个cpu时，选择一组cpu
// 请求8个cpu时，选择两组cpu

func TestResolve(t *testing.T) {
	tests := []struct {
		name  string
		lives []int
		num   int
		want  [][]int
	}{
		{
			name:  "1",
			lives: []int{0, 2, 4, 5, 6, 7},
			num:   2,
			want:  [][]int{{0, 2}},
		},
		{
			name:  "2",
			lives: []int{0, 2, 4, 5, 6, 7},
			num:   1,
			want:  [][]int{{0}, {2}},
		},
		{
			name:  "3",
			lives: []int{0, 2, 4, 5, 6, 7},
			num:   4,
			want:  [][]int{{4, 5, 6, 7}},
		},
		{
			name:  "4",
			lives: []int{0, 4, 5, 6, 7},
			num:   2,
			want:  [][]int{{4, 5}, {4, 6}, {4, 7}, {5, 6}, {5, 7}, {6, 7}},
		},
		{
			name:  "5",
			lives: []int{0, 1, 2, 3, 4, 5, 6, 7},
			num:   8,
			want:  [][]int{{0, 1, 2, 3, 4, 5, 6, 7}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, resolve(test.lives, test.num))
		})
	}
}

func resolve(lives []int, num int) [][]int {
	if num == 8 && len(lives) == 8 {
		return [][]int{lives}
	}
	first := make([]int, 0, 4)
	second := make([]int, 0, 4)
	for _, v := range lives {
		if v < 4 {
			first = append(first, v)
		} else {
			second = append(second, v)
		}
	}
	if len(first)+len(second) < num {
		return nil
	}
	result := make([][]int, 0)
	subFirst, subSecond := len(first)-num, len(second)-num
	if subFirst == subSecond {
		if subFirst > 0 {
			result = append(result, getAll(first, num)...)
			result = append(result, getAll(second, num)...)
		}
		return result
	}
	if (subFirst == 0) || (subFirst == 2 && (subSecond == 1 || subSecond == 3)) || (subFirst == 1 && subSecond == 3) || (subFirst > 0 && subSecond < 0) {
		result = append(result, getAll(first, num)...)
	}
	if (subSecond == 0) || (subSecond == 2 && (subFirst == 1 || subFirst == 3)) || (subSecond == 1 && subFirst == 3) || (subSecond > 0 && subFirst < 0) {
		result = append(result, getAll(second, num)...)
	}
	return result
}

func getAll(arr []int, num int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(start int)
	backtracking = func(start int) {
		if len(path) == num {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(arr); i++ {
			path = append(path, arr[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	fmt.Println("-------", result)
	return result
}
