package stackAndQueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0, len(s)/2)
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')', ']', '}':
			if len(stack) == 0 || (stack[len(stack)-1] != m[v]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "1",
			s:    "()",
			want: true,
		},
		{
			name: "2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "3",
			s:    "(]",
			want: false,
		},
		{
			name: "4",
			s:    "]",
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, isValid(test.s))
		})
	}
}
