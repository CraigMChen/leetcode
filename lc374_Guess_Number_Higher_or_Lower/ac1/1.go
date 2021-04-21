package ac1

func guessNumber(n int) int {
	l, r := 0, 1 << 31 - 1
	for l <= r {
		m := (r - l) >> 1 + l
		g := guess(m)
		if g == 0 {
			return m
		} else if g > 0 {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return 0
}

func guess(num int) int {
	return 0
}
