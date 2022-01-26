package ac1

// 动态规划
// 设 dp[i][j] 为走到坐标(i,j)的路线的数量，则
// 若点 (i,j) 上有障碍，即 obstacleGrid[i][j] == 1，dp[i][j] = 0
// 否则
// dp[0][0] = 1
// dp[i][j] = dp[i-1][j] + dp[i][j-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([]int, len(obstacleGrid[0]))
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if i == 0 && j == 0 {
				dp[j] = 1
			} else if j != 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[len(obstacleGrid[0])-1]
}
