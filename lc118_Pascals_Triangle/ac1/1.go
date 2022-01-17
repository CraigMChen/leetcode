package ac1

//func generate(numRows int) [][]int {
//	triangle := make([][]int, numRows)
//	for i := 0; i < numRows; i++ {
//		triangle[i] = make([]int, i+1)
//		for j := 0; j <= i; j++ {
//			if j == 0 || j == i {
//				triangle[i][j] = 1
//			} else {
//				triangle[i][j] = triangle[i-1][j] + triangle[i-1][j-1]
//			}
//		}
//	}
//	return triangle
//}


// 预处理

var triangle [][]int

func init() {
	triangle = make([][]int, 30)
	for i := 0; i < 30; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				triangle[i][j] = 1
			} else {
				triangle[i][j] = triangle[i-1][j] + triangle[i-1][j-1]
			}
		}
	}
}

func generate(numRows int) [][]int {
	return triangle[:numRows]
}
