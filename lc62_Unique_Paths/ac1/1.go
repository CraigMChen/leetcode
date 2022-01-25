package ac1

// 动态规划
// 设 dp[i][j] 表示走到坐标为(i,j)的点的走法数量，则
// dp[0][0] = 1
// dp[i][j] = dp[i-1][j] + dp[i][j-1]
// dp[m-1][n-1]即为答案
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[j] = 1
			} else if j != 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}
