package ac1

// 原地旋转
// 元素 matrix[i][j] 顺时针旋转 90° 后为 matrix[j][n-1-i]
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n>>1; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] =
				matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		}
	}
}
