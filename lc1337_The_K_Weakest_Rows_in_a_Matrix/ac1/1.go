package ac1

import "sort"

// 通过二分算出矩阵中每一行1的个数，排序
func kWeakestRows(mat [][]int, k int) (ans []int) {
	sums := make([][2]int, len(mat))
	for i := 0; i < len(mat); i++ {
		l, r := 0, len(mat[i])
		for l < r {
			m := (r - l) >> 1 + l
			if mat[i][m] == 1 {
				l = m + 1
			} else {
				r = m
			}
		}
		sums[i][0] = l
		sums[i][1] = i
	}
	sort.Slice(sums, func(i, j int) bool {
		if sums[i][0] == sums[j][0] {
			return sums[i][1] < sums[j][1]
		}
		return sums[i][0] < sums[j][0]
	})

	for i := 0; i < k; i++ {
		ans = append(ans, sums[i][1])
	}
	return ans
}
