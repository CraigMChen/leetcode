package ac1

// 动态规划
// 设 dp[i][j] 表示走到坐标(i,j)的最小和，则
// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
func minPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i != 0 && j != 0 {
				dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
			} else if i != 0 {
				dp[j] = dp[j] + grid[i][j]
			} else if j != 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else {
				dp[j] = grid[i][j]
			}
		}
	}
	return dp[len(grid[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
