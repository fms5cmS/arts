package stackAndQueue

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

// O(N) 时间、空间复杂度
func backspaceCompare(s string, t string) bool {
	return bytes.Equal(handleBackspace(s), handleBackspace(t))
}

func handleBackspace(s string) []byte {
	stack := make([]byte, 0, len(s))
	for _, v := range s {
		if v == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, byte(v))
		}
	}
	return stack
}

// O(N) 时间复杂度
// O(1) 空间复杂度
func backspaceCompare2(s string, t string) bool {
	// 从后往前遍历
	i, j := len(s)-1, len(t)-1
	countS, countT := 0, 0
	for i >= 0 || j >= 0 {
		// 倒序遍历 s 找到 backspace 后的第一个有效字符的索引
		for i >= 0 && (countS > 0 || s[i] == '#') {
			if s[i] == '#' {
				countS++
			} else {
				countS--
			}
			i--
		}
		// 倒序遍历 t
		for j >= 0 && (countT > 0 || t[j] == '#') {
			if t[j] == '#' {
				countT++
			} else {
				countT--
			}
			j--
		}
		// 对字符比较
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i*j <= 0 { // 下面 case4
			return false
		}
		i--
		j--
	}
	return true
}

func TestBackspaceCompare(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{
			name: "1",
			s:    "ab#c",
			t:    "ad#c",
			want: true,
		},
		{
			name: "2",
			s:    "a#c",
			t:    "b",
			want: false,
		},
		{
			name: "3",
			s:    "ab##",
			t:    "c#d#",
			want: true,
		},
		{
			name: "4",
			s:    "bxj##tw",
			t:    "bxj###tw",
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, backspaceCompare(test.s, test.t))
			assert.Equal(t, test.want, backspaceCompare2(test.s, test.t))
		})
	}
}
