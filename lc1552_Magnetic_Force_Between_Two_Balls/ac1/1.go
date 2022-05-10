package ac1

import "sort"

// 二分查找
// 对可能的距离进行二分暴力查找，对于每个可能的结果，可以用O(n)的复杂度求出最多可以放多少个球
func maxDistance(position []int, m int) int {
	sort.Ints(position)

	l, r := 1, position[len(position)-1]+1
	for l < r {
		mid := (r-l)>>1 + l
		if getNum(position, mid) < m {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l - 1
}

func getNum(position []int, d int) int {
	res := 1
	last := position[0]
	for i := 1; i < len(position); i++ {
		if position[i]-last >= d {
			res++
			last = position[i]
		}
	}
	return res
}
