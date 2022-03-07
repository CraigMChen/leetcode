package ac1

// 动态规划
// 设 dp[i] 表示和为i的完全平方数的最小个数，则
// dp[i] = min(dp[i - j * j]) + 1, j * j <= i
func numSquares(n int) int {
	dp := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		tmp := 0x3f3f3f3f
		for j := 1; j * j <= i; j++ {
			tmp = min(tmp, dp[i - j * j])
		}
		dp[i] = tmp + 1
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
