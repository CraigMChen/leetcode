package ac1

// 动态规划
// 设 dp[i][j] 表示 s[0:i] 和 t[0:j] 的最长公共子序列的长度，则
// dp[i][j] = 0 (i == 0 || j == 0)
// dp[i][j] = dp[i-1][j-1] + 1 (s[i-1] == t[j-1])
// dp[i][j] = max(dp[i][j-1], dp[i-1][j]) (s[i-1] != t[j-1])
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	dp[0] = make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
