package ac1

import "sort"

// 以信封的宽（即envelopes[i][0]）为第一关键字升序排序，长（即envelopes[i][1]）为第二关键字降序排序
// 求排序之后信封长的最长上升子序列的长度1
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	var d []int
	d = append(d, -1)
	d = append(d, envelopes[0][1])
	ans := 1
	for i := 0; i < len(envelopes); i++ {
		if envelopes[i][1] > d[ans] {
			d = append(d, envelopes[i][1])
			ans++
		} else {
			l, r := 0, ans
			for l < r {
				m := (r - l) >> 1 + l
				if d[m] < envelopes[i][1] {
					l = m + 1
				} else {
					r = m
				}
			}
			d[l] = envelopes[i][1]
		}
	}
	return ans
}
