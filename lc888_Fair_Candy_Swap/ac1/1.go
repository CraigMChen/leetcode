package ac1

import "sort"

// 二分
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sort.Ints(bobSizes)
	s1, s2 := 0, 0
	for i := range aliceSizes {
		s1 += aliceSizes[i]
	}
	for i := range bobSizes {
		s2 += bobSizes[i]
	}
	x := s1 - s2
	for i := range aliceSizes {
		l, r := 0, len(bobSizes)
		for l < r {
			m := (r - l) >> 1 + l
			tmp := 2 * aliceSizes[i] - 2 * bobSizes[m]
			if tmp == x {
				return []int{aliceSizes[i], bobSizes[m]}
			}
			if tmp > x {
				l = m + 1
			} else {
				r = m
			}
		}
	}
	return nil
}
