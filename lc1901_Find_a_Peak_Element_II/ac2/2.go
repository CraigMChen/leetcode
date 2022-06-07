package ac2

// 模拟
// 随机选一个元素，这里选择[0,0]
// 将该元素与上下左右四个元素相比较，每次向最大的元素移动，直到该中间的元素是顶峰元素
func findPeakGrid(mat [][]int) []int {
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i, j := 0, 0; i < len(mat) && j < len(mat[0]); {
		up, down, left, right := -1, -1, -1, -1
		if i > 0 {
			up = mat[i-1][j]
		}
		if i < len(mat)-1 {
			down = mat[i+1][j]
		}
		if j > 0 {
			left = mat[i][j-1]
		}
		if j < len(mat[0])-1 {
			right = mat[i][j+1]
		}
		max := getMax(getMax(up, down), getMax(left, right))
		if mat[i][j] > max {
			return []int{i, j}
		}
		if max == up {
			i--
		} else if max == down {
			i++
		} else if max == left {
			j--
		} else {
			j++
		}
	}
	return []int{-1, -1}
}
