package ac1

// 贪心
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func maxProfit(prices []int) int {
	res := 0
	least := prices[0]
	for i := 1; i < len(prices); i++ {
		res = max(res, prices[i]-least)
		least = min(least, prices[i])
	}
	return res
}
