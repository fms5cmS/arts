package hash

import (
	"sort"
	"strings"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	m := make(map[string][]string)
	// keys := make([]string, 0)
	for i, str := range strs {
		fields := strings.Split(str, "")
		sort.Strings(fields)
		key := strings.Join(fields, "")
		// if _, exists := m[key]; !exists {
		// 	keys = append(keys, key)
		// }
		m[key] = append(m[key], strs[i])
	}
	// sort.Strings(keys)
	// for _, key := range keys {
	// 	result = append(result, m[key])
	// }
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "2",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "3",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("actual: ", groupAnagrams(test.strs))
			t.Log("want:   ", test.want)
		})
	}
}
