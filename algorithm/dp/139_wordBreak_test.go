package dp

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	// 字符串长度为 i 的话，dp[i] 为 true，表示可以拆分为一个或多个在字典中出现的单词。
	dp := make([]bool, len(s)+1)
	dp[0] = true
	// 如果先遍历物品后遍历背包，需要先把所有可能的字串都保存下来
	// 先遍历背包后遍历物品
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			word := s[j:i]
			if wordSet[word] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
