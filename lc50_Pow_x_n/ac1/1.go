package ac1

// 快速幂
func myPow(x float64, n int) float64 {
	isPositive := true
	if n < 0 {
		n = -n
		isPositive = false
	}
	var ans float64 = 1
	for n > 0 {
		if n & 1 == 1 {
			ans *= x
		}
		x *= x
		n >>= 1
	}
	if !isPositive {
		ans = 1 / ans
	}
	return ans
}
