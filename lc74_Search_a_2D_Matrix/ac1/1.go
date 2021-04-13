package ac1

// 二分
func searchMatrix(matrix [][]int, target int) bool {
	u, d := 0, len(matrix)
	row := -1
	for u < d {
		m := (d - u) >> 1 + u
		if target >= matrix[m][0] && target <= matrix[m][len(matrix[m]) - 1] {
			row = m
			break
		} else if target > matrix[m][len(matrix[m]) - 1] {
			u = m + 1
		} else {
			d = m
		}
	}
	if row == -1 {
		return false
	}

	l, r := 0, len(matrix[row])
	for l < r {
		m := (r - l) >> 1 + l
		if matrix[row][m] == target {
			return true
		} else if matrix[row][m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return false
}
