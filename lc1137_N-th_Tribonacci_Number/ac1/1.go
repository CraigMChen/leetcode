package ac1

func tribonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}
	x, y, z := 0, 1, 1
	m := 0
	for i := 3; i <= n; i++ {
		m = x + y + z
		x = y
		y = z
		z = m
	}
	return m
}
