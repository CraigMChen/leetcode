package ac1

// 二分查找
func minSpeedOnTime(dist []int, hour float64) int {
	getHour := func(m int) float64 {
		hour := 0.0
		for i := 0; i < len(dist); i++ {
			if i == len(dist)-1 {
				hour += float64(dist[i]) / float64(m)
			} else {
				hour += float64((dist[i]-1)/m + 1)
			}
		}
		return hour
	}

	max := 0
	for i := 0; i < len(dist); i++ {
		if dist[i] > max {
			max = dist[i]
		}
	}
	var l, r = 1, max*100 + 1 // 由于hour最多只有2位小数，所以上限是最大值*100
	for l < r {
		m := (r-l)>>1 + l
		if getHour(m) <= hour {
			r = m
		} else {
			l = m + 1
		}
	}
	if l == max*100+1 {
		return -1
	}
	return l
}
