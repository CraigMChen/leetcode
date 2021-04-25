package ac2

import "sort"

// 双指针
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	h := 0
	ans := 0
	for i := 0; i < len(houses); i++ {
		for h < len(heaters) && houses[i] >= heaters[h] {
			h++
		}
		dis := 0
		if h == 0 {
			dis = abs(houses[i] - heaters[h])
		} else if h == len(heaters) {
			dis = abs(houses[i] - heaters[h - 1])
		} else {
			dis = min(abs(houses[i] - heaters[h]), abs(houses[i] - heaters[h - 1]))
		}
		if dis > ans {
			ans = dis
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
