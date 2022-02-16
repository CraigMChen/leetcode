package ac2

// 动态规划
// 设dp[i][j]表示t[i]之后字符j出现的第一个位置，则状态转移方程为
// dp[i][j] = i, t[i] == j
// dp[i][j] = dp[i + 1][j], t[i] != j
//func isSubsequence(s string, t string) bool {
//	dp := make([][]int, len(t))
//	for i := len(t) - 1; i >= 0; i-- {
//		dp[i] = make([]int, 26)
//		for j := 0; j < 26; j++ {
//			if int(t[i]-'a') == j {
//				dp[i][j] = i
//			} else {
//				if i == len(t)-1 {
//					dp[i][j] = -1
//				} else {
//					dp[i][j] = dp[i+1][j]
//				}
//			}
//		}
//	}
//
//	next := 0
//	for i := 0; i < len(s); i++ {
//		if next >= len(t) || dp[next][s[i]-'a'] == -1 {
//			return false
//		}
//		next = dp[next][s[i]-'a'] + 1
//	}
//	return true
//}

// 2022.02.16
func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(t))
	for i := len(t) - 1; i >= 0; i-- {
		dp[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			if t[i] == 'a'+byte(j) {
				dp[i][j] = i
			} else if i == len(t)-1 {
				dp[i][j] = len(t)
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	next := 0
	for i := 0; i < len(s); i++ {
		if next == len(t) || dp[next][s[i]-'a'] == len(t) {
			return false
		}
		next = dp[next][s[i]-'a'] + 1
	}
	return true
}
