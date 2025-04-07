package ac1

// 动态规划
// 根据买卖的次数，可以分为以下 5 种状态：
// 1. 未进行任何买卖
// 2. 买入 1 次，记为 buy1[i]
// 3. 卖出 1 次，记为 sell1
// 4. 买入 2 次，记为 buy2
// 5. 卖出 2 次，记为 sell2
// 考虑每种情况每一天能够获得的最大收益
// 其中状态 1 每天的收益一定是 0，可以忽略
// 剩余四种状态在第 i 天的收益分别记为 buy1[i], sell1[i], buy2[i], sell2[i]
// buy1[i] = -prices[i]
// sell1[i] = max(sell1[i-1], buy1[i] + prices[i])
// buy2[i] = max(buy2[i-1], sell1[i] - prices[i])
// sell2[i] = max(sell1[i-1], buy2[i] + prices[i])
// 最后一天的 sell2 即为最终能获得的最大收益
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}
