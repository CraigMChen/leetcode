package ac1

// 二分
// 从1-1000遍历x，对于每个x，二分搜索找出y3
func findSolution(customFunction func(int, int) int, z int) [][]int {
	var ans [][]int
	for i := 1; i <= 1000; i++ {
		l, r := 1, 1001
		for l < r {
			m := (r - l) >> 1 + l
			if customFunction(i, m) == z {
				ans = append(ans, []int{i, m})
				break
			} else if customFunction(i, m) < z {
				l = m + 1
			} else {
				r = m
			}
		}
	}
	return ans
}
