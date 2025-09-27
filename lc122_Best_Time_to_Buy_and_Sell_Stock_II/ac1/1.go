package ac1

// 贪心
// 求所有递增区间和
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		res += max(0, prices[i]-prices[i-1])
	}
	return res
}
