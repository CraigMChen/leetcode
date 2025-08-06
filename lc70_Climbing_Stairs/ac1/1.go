package ac1

// 动态规划
func climbStairs(n int) int {
	x, y := 1, 1
	for i := 2; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}
