package ac1

// 二分
func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		l, r := 0, n
		for l < r {
			m := (r - l) >> 1 + l
			if grid[i][m] >= 0 {
				l = m + 1
			} else {
				r = m
			}
		}
		ans += n - l
	}
	return ans
}
