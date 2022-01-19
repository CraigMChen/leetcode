package ac1

// 二维前缀和（动态规划思想）
// 题目意思为：对于每个(i,j)，求一个矩阵中所有数字的和，该矩阵的左上角为(i-k,j-k)，右下角为(i+k,j+k)
// 可以用二维前缀和解决。
// 设 preSum[i][j] 为二维数组 A 中，左上角为 (0,0) ，右下角为(i,j) 的矩形中所有数的和，即 A 的二维前缀和
// 根据容斥原理可知，状态转移方程为：
// preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + A[i][j]
// 以行优先的顺序，就可以用 O(1) 的时间内求出每个 (i,j) 的前缀和
//
// 求得前缀和后，就可以用 O(1) 的时间求出二维数组内任意一个矩阵 (x1,y1)(x2,y2) 的和
// 根据容斥原理
// s(x1,y1)(x2,y2) = preSum[x2][y2] - preSum[x1-1][y2] - preSum[x2][y1-1] + preSum[x1-1][y1-1]
func matrixBlockSum(mat [][]int, k int) [][]int {
	preSum := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		preSum[i] = make([]int, len(mat[i]))
		for j := 0; j < len(mat[i]); j++ {
			sa, sb, sc := 0, 0, 0
			if i != 0 {
				sa = preSum[i-1][j]
			}
			if j != 0 {
				sb = preSum[i][j-1]
			}
			if i != 0 && j != 0 {
				sc = preSum[i-1][j-1]
			}
			preSum[i][j] = sa + sb - sc + mat[i][j]
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			x1 := max(0, i-k)
			y1 := max(0, j-k)
			x2 := min(len(mat)-1, i+k)
			y2 := min(len(mat[i])-1, j+k)

			s1, s2, s3 := 0, 0, 0
			if x1 != 0 {
				s1 = preSum[x1-1][y2]
			}
			if y1 != 0 {
				s2 = preSum[x2][y1-1]
			}
			if x1 != 0 && y1 != 0 {
				s3 = preSum[x1-1][y1-1]
			}
			mat[i][j] = preSum[x2][y2] - s1 - s2 + s3
		}
	}
	return mat
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
