package ac1

// 快速乘，二分
func divide(dividend int, divisor int) int {
	var isPositive int64 = 1
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		isPositive = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	var l, r int64 = 0, int64(dividend) + 1
	var ans int64 = 0
	for l < r {
		m := ((r - l) >> 1) + l
		if quickMut(m, int64(divisor)) <= int64(dividend) {
			ans = m
			l = m + 1
		} else {
			r = m
		}
	}
	ans = ans * isPositive
	if ans < -1 << 31 || ans > 1 << 31 - 1 {
		ans = 1 << 31 - 1
	}
	return int(ans)
}

func quickMut(a, b int64) int64 {
	var res int64 = 0
	for b > 0 {
		if b & 1 == 1 {
			res += a
		}
		a += a
		b >>= 1
	}
	return res
}
