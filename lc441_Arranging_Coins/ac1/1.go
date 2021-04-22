package ac1

// 二分
func arrangeCoins(n int) int {
	l, r := 0, n + 1
	ans := 0
	for l < r {
		m := (r - l) >> 1 + l
		if m * (m + 1) / 2 <= n {
			ans = m
			l = m + 1
		} else {
			r = m
		}
	}
	return ans
}
