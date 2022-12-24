package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func plusOne(digits []int) []int {
	carry := 1
	result := make([]int, 0)
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}
	if carry == 1 {
		result = append(result, carry)
	}
	result = append(result, digits...)
	return result
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{
			name:   "1",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			name:   "2",
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2},
		},
		{
			name:   "3",
			digits: []int{9},
			want:   []int{1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, plusOne(test.digits))
		})
	}
}
