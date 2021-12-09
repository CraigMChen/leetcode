package ac2

func maxScoreSightseeingPair(values []int) int {
	res := -0x3f3f3f
	max := values[0]
	for i := 1; i < len(values); i++ {
		res = getMax(res, max+values[i]-i)
		max = getMax(max, values[i]+i)
	}
	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
