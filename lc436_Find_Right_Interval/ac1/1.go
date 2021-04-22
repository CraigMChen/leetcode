package ac1

import "sort"

// 排序，二分查找
func findRightInterval(intervals [][]int) []int {
	mp := make(map[int]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		mp[intervals[i][0]] = i
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		l, r := 0, len(intervals)
		for l < r {
			m := (r-l)>>1 + l
			if intervals[m][0] < intervals[i][1] {
				l = m + 1
			} else {
				r = m
			}
		}
		if l == len(intervals) {
			ans[mp[intervals[i][0]]] = -1
		} else {
			ans[mp[intervals[i][0]]] = mp[intervals[l][0]]
		}
	}
	return ans
}
