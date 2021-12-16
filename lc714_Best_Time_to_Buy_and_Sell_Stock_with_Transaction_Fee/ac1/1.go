package ac1

// 动态规划
// 与lc122类似
func maxProfit(prices []int, fee int) int {
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		x := max(dp0, dp1 + prices[i] - fee)
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
