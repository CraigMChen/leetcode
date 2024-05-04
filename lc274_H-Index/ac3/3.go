package ac3

func hIndex(citations []int) int {
	count := make([]int, 1001)
	for i := 0; i < len(citations); i++ {
		count[citations[i]]++
	}
	h := 0
	for i := len(count) - 1; i >= 0 && i > h; i-- {
		h = min(i, h+count[i])
	}
	return h
}
