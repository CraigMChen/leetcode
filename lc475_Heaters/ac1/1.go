package ac1

import "sort"

// 二分
// 先对heaters进行排序
// 找到每个房子与最近的加热器的距离，其中的最大值即为答案
func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	ans := 0
	for i := 0; i < len(houses); i++ {
		l, r := 0, len(heaters)
		for l < r {
			m := (r - l) >> 1 + l
			if heaters[m] < houses[i] {
				l = m + 1
			} else {
				r = m
			}
		}
		dis := 0
		if l == 0 {
			dis = abs(houses[i] - heaters[l])
		} else if l == len(heaters) {
			dis = abs(houses[i] - heaters[l - 1])
		} else if houses[i] == heaters[l] {
			dis = 0
		} else {
			dis = min(abs(houses[i] - heaters[l]), abs(houses[i] - heaters[l - 1]))
		}
		if dis > ans {
			ans = dis
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
