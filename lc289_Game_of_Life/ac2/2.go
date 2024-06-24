package ac2

// 额外状态
// 加入两种额外状态
// 初始状态 -> 新状态 => 额外状态
// 0        -> 1     => -1
// 1        -> 0     =>  2
// 最后再更新一遍数组
// 空间复杂度为 O(1)
func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := 0
			if i > 0 {
				if board[i-1][j] > 0 {
					sum += 1
				}
				if j > 0 && board[i-1][j-1] > 0 {
					sum += 1
				}
				if j < n-1 && board[i-1][j+1] > 0 {
					sum += 1
				}
			}
			if j > 0 && board[i][j-1] > 0 {
				sum += 1
			}
			if j < n-1 && board[i][j+1] > 0 {
				sum += 1
			}
			if i < m-1 {
				if board[i+1][j] > 0 {
					sum += 1
				}
				if j > 0 && board[i+1][j-1] > 0 {
					sum += 1
				}
				if j < n-1 && board[i+1][j+1] > 0 {
					sum += 1
				}
			}

			res := 0
			if board[i][j] == 0 && sum == 3 || board[i][j] == 1 && sum >= 2 && sum <= 3 {
				res = 1
			}
			if board[i][j] == 0 && res == 1 {
				board[i][j] = -1
			} else if board[i][j] == 1 && res == 0 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
}
