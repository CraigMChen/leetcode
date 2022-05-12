package ac1

// 二分查找
func minDays(bloomDay []int, m int, k int) int {
	getCount := func(d int) int {
		res := 0
		n := 0
		for i := 0; i < len(bloomDay); i++ {
			if bloomDay[i] <= d {
				n++
			} else {
				n = 0
			}
			if n == k {
				res++
				n = 0
			}
		}
		return res
	}

	max := 0
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] > max {
			max = bloomDay[i]
		}
	}
	l, r := 0, max+1
	for l < r {
		mid := (r-l)>>1 + l
		if getCount(mid) >= m {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == max+1 {
		l = -1
	}
	return l
}
