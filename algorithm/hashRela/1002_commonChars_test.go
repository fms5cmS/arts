package hashRela

func commonChars(words []string) []string {
	result := make([]string, 0)
	frequency := make([][]int, 0)
	// 统计每个字符串的词频
	for _, word := range words {
		row := [26]int{}
		for _, s := range word {
			row[s-'a']++
		}
		frequency = append(frequency, row[:])
	}
	// 查找每个字符在所有字符串中出现最少的频次
	// 仅以第一个字符串的频次为基准即可
	for i := 0; i < len(frequency[0]); i++ {
		pre := frequency[0][i]
		for _, freq := range frequency {
			pre = minOf2Ints(pre, freq[i])
		}
		// 还原字符，i 代表了原本字符-'a'，所以这里要再 +'a'
		tmpStr := string(rune(i + 'a'))
		// 按照出现的频次来决定每个字符的频次
		for j := 0; j < pre; j++ {
			result = append(result, tmpStr)
		}
	}
	return result
}

func minOf2Ints(x, y int) int {
	if x < y {
		return x
	}
	return y
}
