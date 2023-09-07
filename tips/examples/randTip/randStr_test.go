package randTip

import (
	"fmt"
	"testing"
)

func TestRandStr(t *testing.T) {
	chars := getChars(2)
	length := 9
	str := RandStrSimple(chars, length)
	t.Log(str)
	idxBits := len(fmt.Sprintf("%b", len(chars)))
	str = RandStrQuick(chars, length, idxBits)
	t.Log(str)
}

func BenchmarkRandStrSimple(b *testing.B) {
	chars := getChars(3)
	length := 25
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandStrSimple(chars, length)
	}
}

func BenchmarkRandStrQuick(b *testing.B) {
	chars := getChars(3)
	length := 25
	idxBits := len(fmt.Sprintf("%b", len(chars)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandStrQuick(chars, length, idxBits)
	}
}

func getChars(tp int8) string {
	switch tp {
	case 0:
		fallthrough
	case 1:
		return "0123456780"
	case 2:
		return "abcdefghijklmnopqrstuvwxyz"
	case 3:
		return "0123456780abcdefghijklmnopqrstuvwxyz"
	default:
		return ""
	}
}
