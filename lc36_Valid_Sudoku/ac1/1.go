package ac1

// 模拟
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		countLine := make([]bool, 10)
		countCol := make([]bool, 10)
		countSq := make([]bool, 10)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if countLine[board[i][j]-'0'] {
					return false
				}
				countLine[board[i][j]-'0'] = true
			}
			if board[j][i] != '.' {
				if countCol[board[j][i]-'0'] {
					return false
				}
				countCol[board[j][i]-'0'] = true
			}
			m := i/3*3 + j/3
			n := i%3*3 + j%3
			if board[m][n] != '.' {
				if countSq[board[m][n]-'0'] {
					return false
				}
				countSq[board[m][n]-'0'] = true
			}
		}
	}
	return true
}
