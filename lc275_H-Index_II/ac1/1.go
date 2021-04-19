package ac1

// 二分
func hIndex(citations []int) int {
	l, r := 0, len(citations)
	for l < r {
		m := (r-l)>>1 + l
		if len(citations)-m <= citations[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return len(citations) - l
}
