package ac1

type NumMatrix struct {
	matrix [][]int
	preSum [][]int
}


// 二维前缀和 动态规划
// 类似lc1314

func Constructor(matrix [][]int) NumMatrix {
	preSum := make([][]int, len(matrix))
	for i := 0;	i < len(matrix); i++ {
		preSum[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			sa, sb, sc := 0, 0, 0
			if i > 0 {
				sa = preSum[i-1][j]
			}
			if j > 0 {
				sb = preSum[i][j-1]
			}
			if i > 0 && j > 0 {
				sc = preSum[i-1][j-1]
			}
			preSum[i][j] = sa + sb - sc + matrix[i][j]
		}
	}
	return NumMatrix{
		matrix: matrix,
		preSum: preSum,
	}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	s1, s2, s3 := 0, 0, 0
	if row1 > 0 {
		s1 = this.preSum[row1-1][col2]
	}
	if col1 > 0 {
		s2 = this.preSum[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		s3 = this.preSum[row1-1][col1-1]
	}
	return this.preSum[row2][col2] - s1 - s2 + s3
}
