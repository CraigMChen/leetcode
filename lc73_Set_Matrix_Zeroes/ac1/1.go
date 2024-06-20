package ac1

// 用第一行的元素来标记每一列是否有 0
// 用第一列的元素来标记每一行是否有 0
// 用 x 标记第一列是否有 0
func setZeroes(matrix [][]int) {
	x := false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			x = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] != 0 {
			continue
		}
		for j := 1; j < len(matrix); j++ {
			matrix[j][i] = 0
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] != 0 {
			continue
		}
		for j := 1; j < len(matrix[0]); j++ {
			matrix[i][j] = 0
		}
	}
	if x {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
