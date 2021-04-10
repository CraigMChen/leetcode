package ac1

func mySqrt(x int) int {
	l, r := 0, 1 << 16
	ans := -1
	for l < r {
		m := (r - l) / 2 + l
		if m != 0 && m > (1 << 31 - 1) / m {
			r = m
		} else if m * m == x {
			return m
		} else if m * m < x {
			ans = m
			l = m + 1
		} else if m * m > x{
			r = m
		}
	}
	return ans
}
