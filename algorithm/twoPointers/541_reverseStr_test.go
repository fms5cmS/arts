package twoPointers

// 当需要固定规律一段一段去处理字符串的时候，要想想在在for循环的表达式上做做文章。
func reverseStr(s string, k int) string {
	str := []byte(s)
	for i := 0; i < len(str); i += 2 * k {
		// If there are less than 2k but greater than or equal to k characters,
		if i+k <= len(str) {
			// then reverse the first k characters and leave the other as original.
			reverseString(str[i : i+k])
			continue
		}
		// If there are fewer than k characters left, reverse all of them
		reverseString(str[i:])
	}
	return string(str)
}
