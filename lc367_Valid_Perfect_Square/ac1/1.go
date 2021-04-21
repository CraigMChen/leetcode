package ac1

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	 l, r := 2, num / 2 + 1
	for l < r {
		m := (r - l) >> 1 + l
		if m == num / m && num % m == 0{
			return true
		} else if m * m < num {
			l = m + 1
		} else {
			r = m
		}
	}
	return false
}
