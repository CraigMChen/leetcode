package ac1

// 动态规划
// 设 dp[i][j] 表示 s[i:j+1] 范围内的最长回文子序列，则
// 若 i == j ， dp[i][j] = 1
// 若 s[i] == s[j] ， dp[i][j] = dp[i+1][j-1] + 2
// 若 s[i] != s[j] ， dp[i][j] = max(dp[i+1][j], dp[i][j-1])
func longestPalindromeSubseq(s string) int {
	dp := make([]int, len(s))
	x := 0 // x 表示 dp[i+1][j-1]
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			tmp := dp[j]
			if i == j {
				dp[j] = 1
			} else if s[i] == s[j] {
				dp[j] = x + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			x = tmp
		}
	}
	return dp[len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
