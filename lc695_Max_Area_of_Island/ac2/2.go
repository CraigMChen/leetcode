package ac2

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 广搜
// 时间复杂度 O(mn)
// 空间复杂度 O(mn)
func maxAreaOfIsland(grid [][]int) int {
	var queue [][2]int
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			cur := 0
			queue = append(queue, [2]int{i, j})
			for len(queue) > 0 {
				loc := queue[0]
				queue = queue[1:]
				x, y := loc[0], loc[1]
				if grid[x][y] == 0 {
					continue
				}
				grid[x][y] = 0
				cur++
				for d := 0; d < len(directions); d++ {
					xx := x + directions[d][0]
					yy := y + directions[d][1]
					if xx < 0 || xx >= len(grid) || yy < 0 || yy >= len(grid[xx]) || grid[xx][yy] == 0 {
						continue
					}
					queue = append(queue, [2]int{xx, yy})
				}
			}
			ans = max(ans, cur)
		}
	}
	return ans
}
