package ac2

// 从左下角元素grid[m-1][0]开始，若该元素大于0则向右走，否则向上走，每次向右走都可以得出该列的负数个数
// 注意边界条件，如果结束时j等于n，说明从第j列开始，后面所有的数都为负数
func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	i, j := m - 1, 0
	ans := 0
	for i >= 0 && j < n {
		if grid[i][j] >= 0 {
			ans += m - 1 - i
			j++
		} else {
			i--
		}
	}
	if j != n {
		ans += (n - j) * (m - 1 - i)
	}
	return ans
}
