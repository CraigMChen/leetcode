package ac1

func isBadVersion(version int) bool {
	return true
}

// 二分
func firstBadVersion(n int) int {
	l, r := 0, n + 1
	for l < r {
		m := (r - l) / 2 + l
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
