package ac1

// 二分查找
// 从 [0, 1001) 范围中，找到最大的满足 h 指数定义的值
// 即，找到第一个不满足 h 指数定义的值 l，l-1 就是结果
func hIndex(citations []int) int {
	counter := func(h int) int {
		count := 0
		for i := 0; i < len(citations); i++ {
			if citations[i] >= h {
				count++
			}
		}
		return count
	}

	l, r := 0, 1001
	for l < r {
		m := (r-l)>>1 + l
		if counter(m) >= m {
			l = m + 1
		} else {
			r = m
		}
	}
	return l - 1
}
