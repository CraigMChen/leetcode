package ac1

import "sort"

// 贪心
// 求删除的最少区间数使剩余区间不重合，可以转化为 选择最多的不重合的区间数
// 先按照右端点从小到达的顺序对数组进行排序
// 维护一个变量 r 表示上一个被选中的区间的右端点
// 若当前区间与上一个被选中的区间不重合
// 则更新 r，并累加到结果中
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res := 1
	r := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= r {
			res++
			r = intervals[i][1]
		}
	}
	return len(intervals) - res
}
