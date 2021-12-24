package hashRela

// 哈希法，因为只会有小写字母，所以可以使用数组来实现哈希
func canConstruct(ransomNote string, magazine string) bool {
	record := [26]int{}
	// 统计 magazine 中每个字符出现的次数
	for _, char := range magazine {
		record[char-'a']++
	}
	for _, char := range ransomNote {
		record[char-'a']--
		if record[char-'a'] < 0 {
			return false
		}
	}
	return true
}
