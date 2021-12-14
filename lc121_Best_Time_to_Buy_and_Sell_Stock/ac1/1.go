package ac1

func maxProfit(prices []int) int {
	res := 0
	min := 10001
	for _, n := range prices {
		res = getMax(res, n - min)
		min = getMin(min, n)
	}
	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}