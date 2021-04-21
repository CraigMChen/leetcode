package ac1

func kthSmallest(matrix [][]int, k int) int {
	l, r := matrix[0][0], matrix[len(matrix) - 1][len(matrix[0]) - 1] + 1
	for l < r {
		m := (r - l) >> 1 + l
		count := 0
		for i, j := len(matrix) - 1, 0; i >= 0 && j < len(matrix[0]); {
			if matrix[i][j] <= m {
				count += i + 1
				j++
			} else {
				i--
			}
		}
		if count < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
