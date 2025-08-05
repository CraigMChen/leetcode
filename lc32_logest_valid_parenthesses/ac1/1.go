package ac1

// 动态规划
// 用 dp[i] 表示以下标 i 结尾的子串的最长有效括号的长度
// 当 s[i-1] == '(' && s[i] == ')' 时，dp[i] = dp[i-2]+2
// 当 s[i-1] == ')' && s[i-dp[i-1]-2] == '(' 时，dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2
			if i >= 2 {
				dp[i] += dp[i-2]
			}
		} else {
			if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 {
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
