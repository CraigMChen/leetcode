package ac1

// 二分查找
// 用 clearDay 数组记录不下雨的天数
// 用 lastRain 记录第i个湖泊上一次下雨的天数
// 遍历rains，若rains[i]!= 0，在clearDay中用二分查找找到第一个大于湖泊rains[i]上次下雨的天数l
// 若找不到则返回空，否则第l天需要清空的湖泊为rains[i]
// 如果有多余的晴天，则随意选择一个湖泊清空，赋为1即可
func avoidFlood(rains []int) []int {
	var (
		sunnyDays   []int
		lastRainDay = make(map[int]int)
		res         = make([]int, len(rains))
	)
	for i := 0; i < len(rains); i++ {
		if rains[i] == 0 {
			sunnyDays = append(sunnyDays, i)
			continue
		}
		res[i] = -1
		if last, ok := lastRainDay[rains[i]]; !ok {
			lastRainDay[rains[i]] = i
		} else {
			l, r := 0, len(sunnyDays)
			for l < r {
				m := (r-l)>>1 + l
				if sunnyDays[m] > last {
					r = m
				} else {
					l = m + 1
				}
			}
			if l >= len(sunnyDays) {
				return nil
			}
			res[sunnyDays[l]] = rains[i]
			sunnyDays = append(sunnyDays[:l], sunnyDays[l+1:]...) // 这一步需要将用过的晴天从数组中移除，复杂度为O(n)
			lastRainDay[rains[i]] = i
		}
	}
	for i := 0; i < len(res); i++ {
		if res[i] == 0 {
			res[i] = 1
		}
	}
	return res
}
