package ac1

// 动态规划
// 设 dp[i][j] 表示以 matrix[i][j] 为右下角的全为'1'的最大正方形的边长，则
// dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1 (matrix[i][j] == '1')
// dp[i][j] = 0 (matrix[i][j] == '0')
// 则结果为最大的dp[i][j]的平方
// 时间复杂度 O(mn)
// 空间复杂度 O(n)
func maximalSquare(matrix [][]byte) int {
	ans := 0
	dp := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		x := 0
		for j := 0; j < len(matrix[i]); j++ {
			tmp := dp[j]
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				dp[j] = int(matrix[i][j] - '0')
			} else {
				dp[j] = min(x, dp[j], dp[j-1]) + 1
			}
			ans = max(ans, dp[j])
			x = tmp
		}
	}
	return ans * ans
}
