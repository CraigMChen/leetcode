package ac1

// 二分搜索速度k
func minEatingSpeed(piles []int, h int) int {
	l, r := 1, int(1e9+1)
	for l < r {
		m := (r - l) >> 1 +l
		hours := 0
		for i := 0; i < len(piles); i++ {
			hours += piles[i] / m
			if piles[i] % m != 0 {
				hours += 1
			}
		}
		if hours > h {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
