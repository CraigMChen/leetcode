package ac1

// 模拟
// 先处理左端点在新区间左端点之前的区间
// 然后处理新区间
// 最后处理左端点在新区间左端点之后的区间
func insert(intervals [][]int, newInterval []int) [][]int {
	var (
		ans [][]int
		i   = 0
	)
	for ; i < len(intervals) && intervals[i][0] <= newInterval[0]; i++ {
		ans = append(ans, intervals[i])
	}
	if len(ans) == 0 {
		ans = append(ans, newInterval)
	} else {
		n := len(ans)
		if newInterval[0] <= ans[n-1][1] && newInterval[1] > ans[n-1][1] {
			ans[n-1][1] = newInterval[1]
		} else if newInterval[0] > ans[n-1][1] {
			ans = append(ans, newInterval)
		}
	}
	for ; i < len(intervals); i++ {
		n := len(ans)
		if intervals[i][0] <= ans[n-1][1] && intervals[i][1] > ans[n-1][1] {
			ans[n-1][1] = intervals[i][1]
		} else if intervals[i][0] > ans[n-1][1] {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}
