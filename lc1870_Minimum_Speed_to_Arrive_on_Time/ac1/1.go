package ac1

// 二分查找
func minSpeedOnTime(dist []int, hour float64) int {
	getTime := func(speed int) float64 {
		res := float64(0)
		for i := 0; i < len(dist); i++ {
			if i < len(dist)-1 {
				res += float64((dist[i]-1)/speed + 1)
			} else {
				res += float64(dist[i]) / float64(speed)
			}
		}
		return res
	}

	var l, r int = 1, 1e7 + 1
	for l < r {
		m := (r-l)/2 + l
		if getTime(m) <= hour {
			r = m
		} else {
			l = m + 1
		}
	}
	if l == 1e7+1 {
		l = -1
	}
	return l
}
