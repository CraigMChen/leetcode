package ac1

// 二分查找
// 用 clearDay 数组记录不下雨的天数
// 用 lastRain 记录第i个湖泊上一次下雨的天数
// 遍历rains，若rains[i]!= 0，在clearDay中用二分查找找到第一个大于湖泊rains[i]上次下雨的天数l
// 若找不到则返回空，否则第l天需要清空的湖泊为rains[i]
// 如果有多余的晴天，则随意选择一个湖泊清空，赋为1即可
func avoidFlood(rains []int) []int {
	var (
		clearDay []int
		res      = make([]int, len(rains))
		lastRain = make(map[int]int)
	)
	for i := 0; i < len(rains); i++ {
		if rains[i] == 0 {
			clearDay = append(clearDay, i+1)
			continue
		}
		res[i] = -1
		last, ok := lastRain[rains[i]]
		if !ok {
			lastRain[rains[i]] = i + 1
		} else {
			l, r := 0, len(clearDay)
			for l < r {
				m := (r-l)>>1 + l
				if clearDay[m] > last {
					r = m
				} else {
					l = m + 1
				}
			}
			if l >= len(clearDay) {
				return nil
			}
			res[clearDay[l]-1] = rains[i]
			clearDay = append(clearDay[:l], clearDay[l+1:]...) // 这一步需要将用过的晴天从数组中移除，复杂度为O(n)
			lastRain[rains[i]] = i + 1
		}
	}
	for i := 0; i < len(res); i++ {
		if res[i] == 0 {
			res[i] = 1
		}
	}
	return res
}
