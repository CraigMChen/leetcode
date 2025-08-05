package ac1

// 二分查找
func mySqrt(x int) int {
	l, r := 0, x+1
	for l < r {
		m := (r-l)>>1 + l
		if m*m <= x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l - 1
}
