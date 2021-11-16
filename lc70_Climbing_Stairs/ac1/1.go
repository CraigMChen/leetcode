package ac1

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	x, y := 1, 2
	z := 0
	for i := 3; i <= n; i++ {
		z = x + y
		x = y
		y = z
	}
	return z
}
