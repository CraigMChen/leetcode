package ac2

// 先水平翻转，再以主对角线翻转
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n>>1; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
