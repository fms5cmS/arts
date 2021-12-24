package stringRela

// 当需要固定规律一段一段去处理字符串的时候，要想想在在for循环的表达式上做做文章。
func reverseStr(s string, k int) string {
	str := []byte(s)
	for i := 0; i < len(str); i += 2 * k {
		if i+k <= len(str) {
			reverseString(str[i : i+k])
			continue
		}
		reverseString(str[i:])
	}
	return string(str)
}
