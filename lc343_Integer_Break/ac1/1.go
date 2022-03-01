package main

// 动态规划
// 设 dp[i] 表示数字 i 可拆分的整数的最大乘积，则
// dp[i] = max( dp[i-j] * j, (i-j)*j ) , j < i
func integerBreak(n int) int {
	dp := make([]int, n + 1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		m := 0
		for j := 1; j < i; j++ {
			m = max(max(m, dp[i-j] * j), (i-j) * j)
		}
		dp[i] = m
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
