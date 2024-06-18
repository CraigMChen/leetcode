package ac1

// 模拟
// 遍历的过程中，行号和列号的变化方式有四种：(0,+1), (+1,0), (0,-1), (-1,0)
// 可以用数组来表示这四种变化方式
// 用 r 表示当前遍历的是第几层
// 当遍历到当前层的四个顶点：
// 右上 (r, n - 1 - r)
// 右下 (m - 1 - r, n - 1 - r)
// 左下 (m - 1 - r, r)
// 左上 (r + 1, r)
// 时，需要变化方向
// 当遍历到当前层的左上顶点时，下一个点进入新的一层
// 防止重复访问，如矩阵 [[1,2,3,4],[5,6,7,8]]，用上面的方法会出现重复访问 1
// 需要统计当前已经遍历的数量 count，当 count == m * n 时停止遍历
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	x := [4]int{0, 1, 0, -1}
	y := [4]int{1, 0, -1, 0}
	r := 0
	dir := 0
	var ans []int
	count := 0
	for i, j := 0, 0; i >= r && i <= m-1-r && j >= r && j <= n-1-r; {
		newRound := false
		ans = append(ans, matrix[i][j])
		count++
		if count == m*n {
			break
		}
		if i == r+1 && j == r && dir == 3 {
			newRound = true
		}
		if i == r && j == n-1-r && dir == 0 || i == m-1-r && j == n-1-r && dir == 1 ||
			i == m-1-r && j == r && dir == 2 || i == r+1 && j == r && dir == 3 {
			dir = (dir + 1) % 4
		}
		i += x[dir]
		j += y[dir]
		if newRound {
			r++
		}
	}
	return ans
}
