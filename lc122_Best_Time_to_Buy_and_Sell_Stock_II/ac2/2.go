package ac2

// 动态规划
// 设 dp[i][0] 表示第i天不持股能获得的最大利润
// dp[i][1] 表示第i天持有股票能获得的最大利润，则
// dp[0][0] = 0, dp[0][1] = -prices[i]
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + price[i])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - price[i])
// dp[len(prices)-1][0] 即为答案
func maxProfit(prices []int) int {
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		x := max(dp0, dp1 + prices[i])
		y := max(dp1, dp0 - prices[i])
		dp0, dp1 = x, y
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
