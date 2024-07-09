package ac1

import "sort"

// 排序
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 || intervals[i][0] > res[len(res)-1][1] {
			res = append(res, []int{intervals[i][0], intervals[i][1]})
		} else if intervals[i][1] > res[len(res)-1][1] {
			res[len(res)-1][1] = intervals[i][1]
		}
	}
	return res
}

func merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	var (
		ans  [][]int
		l, r = intervals[0][0], intervals[0][1]
	)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= r &&
			intervals[i][1] > r {
			r = intervals[i][1]
		} else if intervals[i][0] > r {
			ans = append(ans, []int{l, r})
			l, r = intervals[i][0], intervals[i][1]
		}
	}
	ans = append(ans, []int{l, r})
	return ans
}
