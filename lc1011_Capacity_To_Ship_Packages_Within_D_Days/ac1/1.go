package ac1

// 二分
// 直接二分搜索船的运载能力，对于每个m都遍历一边包裹数组判断天数是否在D以内
func shipWithinDays(weights []int, D int) int {
	l, r := 0, 50000 * 500 + 1
	for l < r {
		m := (r - l) >> 1 + l
		days := 0
		sum := 0
		outOfRange := false
		for i := 0; i < len(weights); i++ {
			if weights[i] > m {
				outOfRange = true
				break
			}
			if sum + weights[i] > m {
				sum = weights[i]
				days++
			} else {
				sum += weights[i]
			}
		}
		if outOfRange {
			l = m + 1
			continue
		}
		if sum > 0 {
			days++
		}
		if days > D {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
