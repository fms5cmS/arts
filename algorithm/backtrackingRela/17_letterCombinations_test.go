package backtrackingRela

import (
	"fmt"
	"testing"
)

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
func letterCombinations(digits string) []string {
	result := make([]string, 0)
	str := ""
	if len(digits) == 0 {
		return result
	}
	// 数字和字母的映射
	letterMap := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi",
		'5': "jkl", '6': "mno", '7': "pqrs",
		'8': "tuv", '9': "wxyz",
	}
	// index 用来取 digits 对应位置的值，即数字
	var backtracking func(index int)
	backtracking = func(index int) {
		if index == len(digits) {
			// 由于 Go 中字符串不可被修改，所以这里每次都会是一个新的字符串，不用像 77、216 那样暂存到临时变量中再添加到 result 中
			result = append(result, str)
			return
		}
		// 获得 digits 对应位置的数字所对应的字符串
		letters := letterMap[digits[index]]
		for _, v := range letters {
			str += string(v)
			backtracking(index + 1)
			str = str[:len(str)-1]
		}
	}
	backtracking(0)
	return result
}

func TestLetterCombinations(t *testing.T) {
	result := letterCombinations("23")
	fmt.Println(result)
}
