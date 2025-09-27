package ac1

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 深搜
// 时间复杂度 O(mn)
// 空间复杂度 O(mn)
func maxAreaOfIsland(grid [][]int) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}
		area := 1
		grid[i][j] = 0
		for d := 0; d < len(directions); d++ {
			area += dfs(i+directions[d][0], j+directions[d][1])
		}
		return area
	}

	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}
