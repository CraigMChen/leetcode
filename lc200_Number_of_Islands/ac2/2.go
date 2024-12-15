package ac2

// 广搜
func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '1' {
				continue
			}
			ans++
			grid[i][j] = 'x'
			queue := [][]int{[]int{i, j}}
			for len(queue) > 0 {
				now := queue[0]
				queue = queue[1:]
				ii, jj := now[0], now[1]
				if ii+1 < len(grid) && grid[ii+1][jj] == '1' {
					grid[ii+1][jj] = 'x'
					queue = append(queue, []int{ii + 1, jj})
				}
				if jj+1 < len(grid[i]) && grid[ii][jj+1] == '1' {
					grid[ii][jj+1] = 'x'
					queue = append(queue, []int{ii, jj + 1})
				}
				if ii > 0 && grid[ii-1][jj] == '1' {
					grid[ii-1][jj] = 'x'
					queue = append(queue, []int{ii - 1, jj})
				}
				if jj > 0 && grid[ii][jj-1] == '1' {
					grid[ii][jj-1] = 'x'
					queue = append(queue, []int{ii, jj - 1})
				}
			}
		}
	}
	return ans
}
