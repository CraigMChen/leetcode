package ac1

// 二分查找
func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for i := 0; i < len(chalk); i++ {
		sum += chalk[i]
		if i != 0 {
			chalk[i] += chalk[i-1]
		}
	}

	k %= sum
	l, r := 0, len(chalk)
	for l < r {
		m := (r-l)>>1 + l
		if chalk[m] <= k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
