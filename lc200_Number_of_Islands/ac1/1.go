package ac1

// 深搜
func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				ans++
				grid[i][j] = 'x'
				dfs(grid, i, j)
			}
		}
	}
	return ans
}

func dfs(grid [][]byte, i, j int) {
	if i > 0 && grid[i-1][j] == '1' {
		grid[i-1][j] = 'x'
		dfs(grid, i-1, j)
	}
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		grid[i+1][j] = 'x'
		dfs(grid, i+1, j)
	}
	if j > 0 && grid[i][j-1] == '1' {
		grid[i][j-1] = 'x'
		dfs(grid, i, j-1)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
		grid[i][j+1] = 'x'
		dfs(grid, i, j+1)
	}
}
