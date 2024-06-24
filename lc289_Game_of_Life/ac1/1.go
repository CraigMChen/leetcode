package ac1

// 辅助数组
// 用一个长度为 n 的数组表示上一行的旧值
// 用一个临时变量表示上一个旧值
// 空间复杂度为 O(n)
func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	curLine := make([]int, n)
	lastLine := make([]int, n)
	for i := 0; i < m; i++ {
		last := 0
		for j := 0; j < n; j++ {
			sum := 0
			if i > 0 {
				sum += lastLine[j]
				if j > 0 {
					sum += lastLine[j-1]
				}
				if j < n-1 {
					sum += lastLine[j+1]
				}
			}
			if j > 0 {
				sum += last
			}
			if j < n-1 {
				sum += board[i][j+1]
			}
			if i < m-1 {
				sum += board[i+1][j]
				if j > 0 {
					sum += board[i+1][j-1]
				}
				if j < n-1 {
					sum += board[i+1][j+1]
				}
			}

			res := 0
			if board[i][j] == 0 && sum == 3 || board[i][j] == 1 && sum >= 2 && sum <= 3 {
				res = 1
			}

			curLine[j] = board[i][j]
			last = board[i][j]
			board[i][j] = res
		}
		for j := 0; j < n; j++ {
			lastLine[j] = curLine[j]
		}
	}
}
