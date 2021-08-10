package greedy

import (
	"fmt"
	"sort"
	"testing"
)

// 局部最优：大饼干喂给胃口大的
func findContentChildren(g []int, s []int) int {
	// 先对两个数组排序
	sort.Ints(g)
	sort.Ints(s)
	// index 代表了饼干数组得到下标，从后往前
	index, result := len(s)-1, 0
	// 从后往前遍历小孩数组，大饼干优先满足胃口大的小孩
	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			result++
			index--
		}
	}
	return result
}

func Test455(t *testing.T) {
	fmt.Println(findContentChildren([]int{1,2,3}, []int{1,1}))
}