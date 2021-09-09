package twoPointer

import (
	"fmt"
	"testing"
)

func reverseWords(s string) string {
	bytesOfStr := []byte(s)
	bytesOfStr = removeExtraSpaces(bytesOfStr)
	reverse(bytesOfStr, 0, len(bytesOfStr)-1)
	for i := 0; i < len(bytesOfStr); i++ {
		j := i
		for j < len(bytesOfStr) && bytesOfStr[j] != ' ' {
			j++
		}
		reverse(bytesOfStr, i, j-1)
		i = j
	}
	return string(bytesOfStr)
}

// 双指针去除字符串中多余的空格，O(n)
// 注意：要么参数这里传切片的指针，要么像这样用返回值接受
// 切片实际为结构体，虽然原始切片和传入后的切片指向的底层数组相同，且这里修改了底层数组的值，不过原始切片的长度和容量不变，所以如果不返回值的话，值是错误的
// 注意与 145 中的区别！
func removeExtraSpaces(bytesOfStr []byte) []byte {
	slow, fast := 0, 0
	// 去除字符串前面的空格，fast 会指向第一个非空格的字符
	for len(bytesOfStr) > 0 && fast < len(bytesOfStr) && bytesOfStr[fast] == ' ' {
		fast++
	}
	for ; fast < len(bytesOfStr); fast++ {
		// 去除字符串中间冗余的空格
		if fast-1 > 0 && bytesOfStr[fast-1] == bytesOfStr[fast] && bytesOfStr[fast] == ' ' {
			continue
		} else {
			bytesOfStr[slow] = bytesOfStr[fast]
			slow++
		}
	}
	// 去除字符串末尾的空格
	if slow-1 > 0 && bytesOfStr[slow-1] == ' ' {
		bytesOfStr = bytesOfStr[:slow-1]
	} else {
		bytesOfStr = bytesOfStr[:slow]
	}
	return bytesOfStr
}

// 反转指定范围内的字符串
func reverse(bytesOfStr []byte, start, end int) {
	for start < end {
		bytesOfStr[start], bytesOfStr[end] = bytesOfStr[end], bytesOfStr[start]
		start++
		end--
	}
}

func TestRemoveWords(t *testing.T) {
	s1 := "the sky is blue"
	fmt.Println(reverseWords(s1))
	s2 := "  hello world  "
	fmt.Println(reverseWords(s2))
}