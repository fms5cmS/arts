package twoPointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{
			s:    []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			s:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, test := range tests {
		reverseString(test.s)
		assert.Equal(t, test.want, test.s)
	}
}
