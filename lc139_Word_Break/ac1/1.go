package ac1

// 动态规划
// 设 dp[i] 表示 s 的前 i 个字符（即s[0:i-1]）能否被完整分割
// 对于一个i，可以将s[0:i-1]分割成两个部分
// s[0:j], s[j:i]
// 只要s[0:j]和s[j:i]都在字典中出现过，那么s[i]就可以被完整分割，则
// dp[i] = dp[j] && dict(s[j:i]) (j < i)
// 设空字符串为可完整分割的，即 dp[0] = true
func wordBreak(s string, wordDict []string) bool {
	dictionary := make(map[string]struct{})
	for i := range wordDict {
		dictionary[wordDict[i]] = struct{}{}
	}

	b := []byte(s)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		canBreak := false
		for j := 0; j < i; j++ {
			if _, ok := dictionary[string(b[j:i])]; ok && dp[j]	{
				canBreak = true
				break
			}
		}
		dp[i] = canBreak
	}
	return dp[len(s)]
}
