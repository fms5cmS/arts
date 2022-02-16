package showMeBug

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// CountDuplicateWord 计算相邻重复单词数
// 实现一个函数，计算一个字符串中相邻重复单词的个数(不区分大小写)。两个或两个以上相等的单词依次出现视为一个。
// "dog cat"                  -->  0
// "dog DOG cat"              -->  1
// "apple dog cat"            -->  0
// "pineapple apple dog cat"  -->  0
// "apple     apple dog cat"  -->  1
// "apple dog apple dog cat"  -->  0
// "dog dog DOG dog dog dog"  -->  1
// "dog dog dog dog cat cat"  -->  2
// "cat cat dog dog cat cat"  -->  3
func CountDuplicateWord(s string) int {
	strs := strings.Split(s, " ")
	if len(strs) < 2 {
		return 0
	}
	first, next := 0, 1
	count := 0
	equal := func(src, dst string) bool {
		return strings.ToLower(strings.TrimSpace(src)) == strings.ToLower(strings.TrimSpace(dst))
	}
	for ; next < len(strs); next++ {
		if !equal(strs[first], strs[next]) {
			first = next
			continue
		}
		if next-first == 1 {
			count++
		}
	}
	return count
}

func TestCountDuplicateWord(t *testing.T) {
	assert := assert.New(t)
	tests := []struct{
		input string
		output int
	}{
		{input: "dog cat", output: 0},
		{input: "dog DOG cat"            , output: 1},
		{input: "apple dog cat"          , output: 0},
		{input: "pineapple apple dog cat", output: 0},
		{input: "apple     apple dog cat", output: 1},
		{input: "apple dog apple dog cat", output: 0},
		{input: "dog dog DOG dog dog dog", output: 1},
		{input: "dog dog dog dog cat cat", output: 2},
		{input: "cat cat dog dog cat cat", output: 3},
	}
	for _, test := range tests {
		real := CountDuplicateWord(test.input)
		assert.Equal(real, test.output)
	}
}