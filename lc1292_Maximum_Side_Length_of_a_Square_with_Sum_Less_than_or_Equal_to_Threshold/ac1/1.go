package ac1

// 二维前缀和+二分查找
// 首先求出矩阵mat的二维前缀和preSum
// 然后在[1,min(m,n)]中用二分查找的方法搜索边长k，找到第一个下标l，使得不存在符合条件的正方形
// 则 l-1 即为答案
func maxSideLength(mat [][]int, threshold int) int {
	preSum := make([][]int, len(mat)+1)
	preSum[0] = make([]int, len(mat[0])+1)
	for i := 1; i <= len(mat); i++ {
		preSum[i] = make([]int, len(mat[i-1])+1)
		for j := 1; j <= len(mat[i-1]); j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + mat[i-1][j-1]
		}
	}

	isExisted := func(k int) bool {
		for i := 0; i < len(mat)-k+1; i++ {
			for j := 0; j < len(mat[i])-k+1; j++ {
				if preSum[i+k][j+k]-preSum[i][j+k]-preSum[i+k][j]+preSum[i][j] <= threshold {
					return true
				}
			}
		}
		return false
	}

	l, r := 1, min(len(mat), len(mat[0]))+1
	for l < r {
		m := (r-l)>>1 + l
		if !isExisted(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
