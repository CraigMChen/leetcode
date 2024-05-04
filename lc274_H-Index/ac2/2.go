package ac2

import "sort"

// 快排
func hIndex(citations []int) int {
	sort.Ints(citations)
	h := 0
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return h
}
