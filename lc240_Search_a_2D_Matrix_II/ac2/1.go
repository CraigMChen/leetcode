package ac2

// 把二维数组近似看做一颗二叉搜索树，左下角的元素为树的根节点
func searchMatrix(matrix [][]int, target int) bool {
	i, j := len(matrix) - 1, 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			j++
		} else {
			i--
		}
	}
	return false
}
