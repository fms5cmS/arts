package twoPointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 整体反转+局部反转 ===>  左旋转（单词反转）
// 1. 移除多余的空格
// 2. 整体反转
// 3. 局部反转
func reverseWords(s string) string {
	bytesOfStr := []byte(s)
	bytesOfStr = removeExtraSpace(bytesOfStr)
	// 反转所有的字符
	// "the sky is blue"  --->   "eulb si yks eht"
	reverseRange(bytesOfStr, 0, len(bytesOfStr)-1)
	left, right := 0, 0
	for left < len(bytesOfStr) {
		// 根据空格找到每个单词的范围
		// 注意，这里 for 循环之后，right 指向的是一个空字符串！！所以后面单词反转后需要将 left、right 都指向下一个单词的第一个字符
		for right < len(bytesOfStr) && bytesOfStr[right] != ' ' {
			right++
		}
		// 反转单词，使其变成原始单词
		reverseRange(bytesOfStr, left, right-1)
		right++
		left = right
	}
	return string(bytesOfStr)
}

// 反转指定范围内的字符串
func reverseRange(bytesOfStr []byte, start, end int) {
	for start < end {
		bytesOfStr[start], bytesOfStr[end] = bytesOfStr[end], bytesOfStr[start]
		start++
		end--
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			s:    "a good   example",
			want: "example good a",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, reverseWords(test.s))
	}
}

// 双指针去除字符串中多余的空格，O(n)
// 注意，切片底层实际为结构体，由于这里入参宾并不是指针，所以如果不返回值的话，其实仅改变了底层数组，但是原切片的总长度是没有变化的！
// 注意与 145 中的区别！
func removeExtraSpace(bytesOfStr []byte) []byte {
	slow, fast := 0, 0
	// 找到第一个不为空格的字符
	for fast < len(bytesOfStr) && bytesOfStr[fast] == ' ' {
		fast++
	}
	for ; fast < len(bytesOfStr); fast++ {
		// 连续的空格仅保留一个，所以如果是末尾有多个空格的话，也会保留一个，需要在最后额外判断一下
		if bytesOfStr[fast] == ' ' && bytesOfStr[fast-1] == ' ' {
			continue
		}
		bytesOfStr[slow] = bytesOfStr[fast]
		slow++
	}
	// 如果原字符串最后有多个空格，那么上面处理后最后还会有一个空格，需要移除
	// 注意这里是 slow-1
	if bytesOfStr[slow-1] == ' ' {
		slow--
	}
	return bytesOfStr[:slow]
}

func TestRemoveExtraSpace(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{
			src:  "the sky is blue",
			want: "the sky is blue",
		},
		{
			src:  "  hello world  ",
			want: "hello world",
		},
		{
			src:  "a good   example",
			want: "a good example",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, string(removeExtraSpace([]byte(test.src))))
	}
}
