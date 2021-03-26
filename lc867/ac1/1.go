package ac1

// 申请一块与原矩阵大小相同的内存
func transpose(matrix [][]int) [][]int {
	resMatrix := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		resMatrix[i] = make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			resMatrix[i][j] = matrix[j][i]
		}
	}
	return resMatrix
}
