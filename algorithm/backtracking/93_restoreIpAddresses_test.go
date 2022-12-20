package backtracking

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return nil
	}
	result := make([]string, 0)
	path := make([]string, 0)
	var backtracking func(start int)
	backtracking = func(start int) {
		if len(path) == 4 && start == len(s) {
			str := strings.Join(path, ".")
			result = append(result, str)
		}
		for i := start; i < len(s); i++ {
			if isValidIPSegment(s[start : i+1]) {
				path = append(path, s[start:i+1])
			} else {
				continue
			}
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return result
}

func isValidIPSegment(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	i, _ := strconv.Atoi(s)
	if i < 0 || i > 255 {
		return false
	}
	return true
}

func restoreIpAddresses2(s string) []string {
	result := make([]string, 0)
	// 判断字符串 str 在 [start, end] 区间组成的数字是否合法
	isValid := func(str string) bool {
		if str == "" || (len(str) > 1 && (str[0] == '0' || str[0] == '+')) {
			return false
		}
		after, err := strconv.Atoi(str)
		// 转换失败说明原字符串有特殊字符（注意，这里区分不了 + 和 - 两种特殊字符，所以需要单独判断）
		if err != nil {
			return false
		}
		// 负数相当于判断了首字符为 - 的特殊字符
		if after < 0 || after > 255 {
			return false
		}
		return true
	}
	// 回溯函数
	var backtracking func(str string, validNums []string, startIndex int)
	backtracking = func(str string, validNums []string, startIndex int) {
		if len(validNums) == 4 && startIndex == len(str) {
			// 这里并不是直接将 validNums 添加到返回结果中，所以不需要定义临时变量将 validNums 复制到里面再处理
			result = append(result, strings.Join(validNums, "."))
		}
		for i := startIndex; i < len(str); i++ {
			tempStr := str[startIndex : i+1]
			if !isValid(tempStr) {
				break
			}
			validNums = append(validNums, tempStr)
			backtracking(str, validNums, i+1)
			validNums = validNums[:len(validNums)-1]
		}
	}
	// IP 地址最长为 12 个字符串
	if len(s) > 12 {
		return result
	}
	validNums := make([]string, 0)
	backtracking(s, validNums, 0)
	return result
}

func TestRestoreIpAddress(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output []string
	}{
		{name: "1", input: "25525511135", output: []string{"255.255.11.135", "255.255.111.35"}},
		{name: "2", input: "0000", output: []string{"0.0.0.0"}},
		{name: "3", input: "1111", output: []string{"1.1.1.1"}},
		{name: "4", input: "010010", output: []string{"0.10.0.10", "0.100.1.0"}},
		{name: "5", input: "101023", output: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, restoreIpAddresses(test.input))
			assert.Equal(t, test.output, restoreIpAddresses2(test.input))
		})
	}
}
