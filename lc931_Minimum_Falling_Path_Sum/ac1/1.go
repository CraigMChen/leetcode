package ac1

// 动态规划
// 设 dp[i][j] 表示以 matrix[i][j] 结尾的路径最小和，则
// dp[i][j] = matrix[i][j] + min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1)
// dp[0][j] = matrix[0][j]
// min(dp[len(matrix)-1]) 即为答案
func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 1 {
		return matrix[0][0]
	}
	res := 0x3f3f3f3f
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j+1])
			} else if j == len(matrix[i]) - 1 {
				matrix[i][j] += min(matrix[i-1][j-1], matrix[i-1][j])
			} else {
				matrix[i][j] += min(min(matrix[i-1][j-1], matrix[i-1][j]), matrix[i-1][j+1])
			}
			if i == len(matrix) - 1 {
				res = min(res, matrix[i][j])
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
