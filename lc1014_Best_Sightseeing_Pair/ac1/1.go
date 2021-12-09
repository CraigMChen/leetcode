package ac1

// 表达式 values[i] + values[j] + i - j 可以分为两个部分
// values[i] + i 和 values[j] - j
// 遍历i，对于每个i，前半部分是固定的，后半部分可以通过预处理求出
// 这样的时间复杂度为O(n)，空间复杂为O(n)
// 实际上可以遍历j，并同时维护values[i]+i的最大值
// 这样空间复杂度就可以降到O(1)，见ac2
func maxScoreSightseeingPair(values []int) int {
	maxSuf := make([]int, len(values))
	max := -0x3f3f3f3f
	for i := len(values) - 1; i >= 0; i-- {
		max = getMax(max, values[i] - i)
		maxSuf[i] = max
	}

	res := -0x3f3f3f3f
	for i := 0; i < len(values) - 1; i++ {
		res = getMax(res, values[i]+i+maxSuf[i+1])
	}
	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
