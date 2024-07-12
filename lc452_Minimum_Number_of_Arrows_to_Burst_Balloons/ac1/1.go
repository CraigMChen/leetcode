package ac1

import "sort"

// 排序
// 先将 points 按照左端点从小到大排序
// 遍历 points，计算当前区间与之前区间的交集 [l,r]
// 若当前区间与 [l,r] 无交集，则箭的总数+1
// 否有交集，则该区间不需要额外的箭，更新交集
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ans := 1
	l, r := points[0][0], points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] >= l && points[i][0] <= r {
			l = max(l, points[i][0])
			r = min(r, points[i][1])
		} else {
			ans++
			l = points[i][0]
			r = points[i][1]
		}
	}
	return ans
}
