package ac1

// 二分查找
// 二分查找列col
// 对于某一列col，找出该列的最大元素mat[col][j]
// 若mat[col][j]>mat[col][j-1]且mat[col][j]>mat[col][j+1]，则(col,j)即为答案
// 若左边较大，则搜索左半边的矩阵，否则搜右半边
func findPeakGrid(mat [][]int) []int {
	getMaxRowID := func(m int) int {
		max := -1
		id := -1
		for i := 0; i < len(mat); i++ {
			if mat[i][m] > max {
				max = mat[i][m]
				id = i
			}
		}
		return id
	}

	l, r := 0, len(mat[0])
	for l < r {
		m := (r-l)>>1 + l
		row := getMaxRowID(m)
		if m == 0 && mat[row][m] > mat[row][m+1] || m == len(mat[0])-1 && mat[row][m] > mat[row][m-1] ||
			mat[row][m] > mat[row][m-1] && mat[row][m] > mat[row][m+1] {
			return []int{row, m}
		} else if m == len(mat[0])-1 || mat[row][m] < mat[row][m-1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return []int{getMaxRowID(l), l}
}
