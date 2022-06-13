package ac1

// 二分
// 利用容斥原理可以在O(1)的时间里求出[1, x]内含有因子a或b或c的数（即丑数）的个数：
// num = x / a + x / b + x / c - x / lcm(a, b) - x / lcm(a, c) - lcm(b, c) + lcm(a, b, c)
// 则可以用二分搜索求出满足在区间[1,x]内丑数个数大于等于n的第一个数x，即为答案
func nthUglyNumber(n int, a int, b int, c int) int {
	var gcd, lcm func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}
	lcm = func(a, b int) int {
		return a * b / gcd(a, b)
	}

	getK := func(m int) int {
		return m/a + m/b + m/c - m/lcm(a, b) - m/lcm(a, c) - m/lcm(b, c) + m/lcm(lcm(a, b), c)
	}
	var l, r int = 1, 2e9 + 1
	for l < r {
		m := (r-l)>>1 + l
		if getK(m) >= n {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
