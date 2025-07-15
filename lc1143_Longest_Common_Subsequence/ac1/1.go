package ac1

// 动态规划
// 设 dp[i][j] 表示 s[0:i] 和 t[0:j] 的最长公共子序列的长度，则
// dp[i][j] = 0 (i == 0 || j == 0)
// dp[i][j] = dp[i-1][j-1] + 1 (s[i-1] == t[j-1])
// dp[i][j] = max(dp[i][j-1], dp[i-1][j]) (s[i-1] != t[j-1])
// 可以用滚动数组优化空间
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		x := 0
		for j := 1; j <= len(text2); j++ {
			tmp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = x + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			x = tmp
		}
	}
	return dp[len(text2)]
}
