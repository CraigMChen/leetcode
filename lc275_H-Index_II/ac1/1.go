package ac1

// 二分
// 用二分查找在范围[1, len(citations)]中找到第一个符合条件的l，使得
// citations[len(citations)-l] < l
// 则结果为 l-1
func hIndex(citations []int) int {
	l, r := 1, len(citations)+1
	for l < r {
		m := (r-l)>>1 + l
		if citations[len(citations)-m] < m {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}
