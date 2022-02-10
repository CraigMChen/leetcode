package ac1

// 动态规划
// 设 dp[i][j] 表示以 matrix[i][j] 为右下角的全为'1'的最大正方形的边长，则
// dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1 (matrix[i][j] == '1')
// dp[i][j] = 0 (matrix[i][j] == '0')
// 则结果为最大的dp[i][j]的平方
func maximalSquare(matrix [][]byte) int {
	dp := make([]int, len(matrix[0]))
	x := 0
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			temp := dp[j]
			y := 0
			if j != 0 {
				y = dp[j-1]
			}
			if matrix[i][j] == '1' {
				dp[j] = min(min(x, y), dp[j]) + 1
			} else {
				dp[j] = 0
			}
			x = temp
			if dp[j] > res {
				res = dp[j]
			}
		}
	}
	return res * res
}

func min(a, b int) int{
	if a > b {
		return b
	}
	return a
}
