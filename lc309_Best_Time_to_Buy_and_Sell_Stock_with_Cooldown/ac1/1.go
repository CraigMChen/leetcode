package ac1

// 动态规划
// 设 dp[i][0] 表示第i天持有股票时能获得的最大利润
// dp[i][1] 表示第i天不持股，且第i天结束时进入冷冻期时能获得的最大利润
// dp[i][2] 表示第i天不持股，且第i天结束时不进入冷冻期时能获得的最大利润，则
// dp[i][0] = max(dp[i-1][0], dp[i-1][2] - prices[i])
// dp[i][1] = dp[i-1][0] + prices[i]
// dp[i][2] = max(dp[i-1][1], dp[i-1][2])
// dp[0][0] = -prices[0], dp[0][1] = 0, dp[0][2] = 0
func maxProfit(prices []int) int {
	dp0, dp1, dp2 := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		x := max(dp0, dp2 - prices[i])
		y := dp0 + prices[i]
		z := max(dp1, dp2)
		dp0, dp1, dp2 = x, y, z
	}
	return max(dp1, dp2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
