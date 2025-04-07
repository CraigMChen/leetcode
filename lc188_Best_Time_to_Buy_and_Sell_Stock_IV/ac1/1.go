package ac1

// 动态规划
// 思路与 lc123 类似
// 根据买卖的次数，可以分为以下 2k 种状态：
// 1. 买入 1 次
// 2. 卖出 1 次
// 3. 买入 2 次
// 4. 卖出 2 次
// ...
// 2k-1. 买入 k 次
// 2k. 卖出 k 次
// 考虑每种状态每天的最大收益
// 记买入 i 次在第 j 天的最大收益为 buy[i][j]
// 卖出 i 次在第 j 天的最大收益为 sell[i][j]
// 初始状态：
// buy[i][0] = -prices[0]
// sell[i][0] = 0
// 状态转移方程：
// buy[i][j] = max(buy[i][j-1], sell[i-1][j])
// sell[i][j] = max(sell[i][j-1], buy[i][j])
// 最后一天的 sell[k] 即为最大收益
func maxProfit(k int, prices []int) int {
	buy, sell := make([]int, k+1), make([]int, k+1)
	for i := 1; i <= k; i++ {
		buy[i] = -prices[0]
		sell[i] = 0
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	return sell[k]
}
