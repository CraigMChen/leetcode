package ac1

func getRow(rowIndex int) []int {
	triangle := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		for j := i; j >=0; j-- {
			if j == 0 || j == i {
				triangle[j] = 1
			} else {
				triangle[j] = triangle[j] + triangle[j-1]
			}
		}
	}
	return triangle
}
