package ac3

// 数学
// 所有数字都会落在循环 4→16→37→58→89→145→42→20→4 上或最终变成 1
func isHappy(n int) bool {
	set := map[int]struct{}{
		4:   {},
		16:  {},
		37:  {},
		58:  {},
		89:  {},
		145: {},
		42:  {},
		20:  {},
	}
	for _, ok := set[n]; !ok && n != 1; _, ok = set[n] {
		n = squareSum(n)
	}
	return n == 1
}

func squareSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
