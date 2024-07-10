package ac2

// 模拟
// 遍历 intervals，对于与 newInterval 无交集的区间可以直接插入
// 若区间与新区间有交集，则该区间与 newInterval 合并成一个新区间
func insert(intervals [][]int, newInterval []int) [][]int {
	var (
		ans      [][]int
		l, r     = newInterval[0], newInterval[1]
		inserted = false
	)
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > r {
			if !inserted {
				ans = append(ans, []int{l, r})
				inserted = true
			}
			ans = append(ans, intervals[i])
		} else if intervals[i][1] < l {
			ans = append(ans, intervals[i])
		} else {
			l = min(l, intervals[i][0])
			r = max(r, intervals[i][1])
		}
	}
	if !inserted {
		ans = append(ans, []int{l, r})
	}
	return ans
}
