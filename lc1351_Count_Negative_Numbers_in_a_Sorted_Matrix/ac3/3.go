package ac3

// 分治
// count(u,d,l,r)表示[u,d)行中列号为[l,r)的元素中负数的个数
// mid表示[u,d)的中间行号
// k表示[u,d)的中间行的负数的个数
// 由矩阵的特性可知，负数个数从上到下非递减
// 则count(u,d,l,r) = n - k + count(u, mid, k, r) + count(mid + 1, d, l, k + 1)
func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var count func(u, d, l, r int) int
	count = func(u, d, l, r int) int {
		if d <= u {
			return 0
		}
		k := -1
		mid := (d - u) >> 1 + u
		for i := l; i < r; i++ {
			if grid[mid][i] < 0 {
				k = i
				break
			}
		}
		if k != -1 {
			return n - k + count(u, mid, k, r) + count(mid + 1, d, l, k + 1)
		} else {
			return count(mid + 1, d, l, r)
		}
	}
	return count(0, m, 0, n)
}
