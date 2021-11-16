package ac1

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	x, y := 0, 1
	z := 0
	for i := 2; i <= n; i++ {
		z = x + y
		x = y
		y = z
	}
	return z
}
