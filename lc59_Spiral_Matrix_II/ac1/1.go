package ac1

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	i, j, k := 0, 0, 0
	for x := 1; x <= n*n; x++ {
		res[i][j] = x
		if i == k && j < n-k-1 {
			j++
		} else if j == n-k-1 && i < n-k-1 {
			i++
		} else if i == n-k-1 && j > k {
			j--
		} else if j == k && i > k+1 {
			i--
		} else {
			j++
			k++
		}
	}
	return res
}
