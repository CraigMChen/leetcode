package ac1

// 求递增区间和
func maxProfit(prices []int) int {
	res := 0
	last := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			res += prices[i-1] - last
			last = prices[i]
		}
	}
	return prices[len(prices)-1] - last + res
}
