package backtrackingRela

import (
	"github.com/stretchr/testify/assert"
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
				subStr := s[startIndex : i+1] // 注意这里是 i+1
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
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{
			name: "1",
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "2",
			s:    "a",
			want: [][]string{{"a"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, partition(test.s))
		})
	}
}
