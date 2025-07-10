package ac1

// 模拟
// 遍历的过程中，行号和列号的变化方式有四种：(0,+1), (+1,0), (0,-1), (-1,0)
// 可以用数组来表示这四种变化方式
// 用 cycle 表示当前遍历的是第几层
// 当遍历到当前层的四个顶点：
// 右上 (cycle, col-cycle-1) 且方向向右
// 右下 (row-cycle-1, col-cycle-1)，且方向向下
// 左下 (row-cycle-1, cycle)，且方向向左
// 左上 (cycle+1, cycle)，且方向向上
// 时，需要变化方向
// 当遍历到当前层的左上顶点时，下一个点进入新的一层
// 防止重复访问，如矩阵 [[1,2,3,4],[5,6,7,8]]，用上面的方法会出现重复访问 1
// 需要统计当前已经遍历的数量 k，当 k == row*cycle 时停止遍历
func spiralOrder(matrix [][]int) []int {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col := len(matrix), len(matrix[0])
	total := row * col
	res := make([]int, total)
	k := 0
	i, j, d, cycle := 0, 0, 0, 0
	for k < total {
		res[k] = matrix[i][j]
		if d == 0 && i == cycle && j == col-cycle-1 ||
			d == 1 && i == row-cycle-1 && j == col-cycle-1 ||
			d == 2 && i == row-cycle-1 && j == cycle ||
			d == 3 && i == cycle+1 && j == cycle {
			d = (d + 1) % len(directions)
			if i == cycle+1 && j == cycle {
				cycle++
			}
		}
		i = i + directions[d][0]
		j = j + directions[d][1]
		k++
	}
	return res
}
