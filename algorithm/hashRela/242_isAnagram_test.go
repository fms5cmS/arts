package hashRela

func isAnagram(s string, t string) bool {
	var record [26]int
	// 记录 s 中每个字符出现的次数
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
