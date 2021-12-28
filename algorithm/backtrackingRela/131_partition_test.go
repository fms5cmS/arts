package backtrackingRela

import (
	"fmt"
	"testing"
)

func partition(s string) [][]string {
	result := make([][]string, 0)
	path := make([]string, 0)
	// 判断 str 的字串（指定起始和结束索引）是否是回文串
	isPalindrome := func(str string, start, end int) bool {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
		return true
	}

	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		if startIndex >= len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := startIndex; i < len(s); i++ {
			if isPalindrome(s, startIndex, i) {
				subStr := s[startIndex:i+1] // 注意这里是 i+1
				path = append(path, subStr)
			} else {
				continue
			}
			backtracking(i + 1)
			// 代码走到这里，代表一定往 path 中添加了元素，所以不会存在 len(path) = 0 导致越界的情况
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return result
}

func TestPartition(t *testing.T) {
	src := "aab"
	result := partition(src)
	fmt.Println(result)
}
