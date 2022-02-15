package ac1

// 动态规划
// 设dp[i][j]表示s[i,j]为回文串，则状态转移方程为
// dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])
// 边际条件为
// dp[i][i] = true
// dp[i][i+1] = s[i] == s[i+1]
//func longestPalindrome(s string) string {
//	l := len(s)
//	if l == 0 {
//		return ""
//	}
//	dp := make([][]bool, l)
//	for i := 0; i < l; i++ {
//		dp[i] = make([]bool, l)
//	}
//	palindromeLen := 1
//	start, end := 0, 0
//	for k := 0; k < l; k++ {
//		for i := 0; i+k < l; i++ {
//			j := i + k
//			if i == j {
//				dp[i][j] = true
//			} else if j == i+1 {
//				dp[i][j] = s[i] == s[j]
//			} else {
//				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
//			}
//			if dp[i][j] && j-i+1 > palindromeLen {
//				palindromeLen = j - i + 1
//				start = i
//				end = j
//			}
//		}
//	}
//	return s[start : end+1]
//}

// 2022.02.14
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	l := 1
	start, end := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = true
			} else if j == i + 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] && j - i + 1 > l {
				l = j - i + 1
				start = i
				end = j
			}
		}
	}
	return s[start:end+1]
}