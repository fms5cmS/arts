package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isAnagram(s string, t string) bool {
	var record [26]int
	// record 记录 s 中每个字符出现的次数
	for _, ss := range s {
		record[ss-'a']++
	}
	for _, ts := range t {
		record[ts-'a']--
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	// record 所有元素值均为 0，说明 s 和 t 是字母异位词
	return true
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		result bool
	}{
		{
			s:      "anagram",
			t:      "nagaram",
			result: true,
		},
		{
			s:      "rat",
			t:      "car",
			result: false,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.result, isAnagram(test.s, test.t))
	}
}
