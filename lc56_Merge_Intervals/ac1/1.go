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
