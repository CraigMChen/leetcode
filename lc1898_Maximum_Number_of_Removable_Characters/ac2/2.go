package ac2

// 二分查找
// 用双指针法判断是否为子序列，类似lc392
func maximumRemovals(s string, p string, removable []int) int {
	removed := make([]int, len(s))
	for i := 0; i < len(removed); i++ {
		removed[i] = 1e5 + 1
	}
	for i := 0; i < len(removable); i++ {
		removed[removable[i]] = i
	}

	isSubStr := func(k int) bool {
		i, j := 0, 0
		for i < len(s) && j < len(p) {
			if removed[i] >= k && s[i] == p[j] {
				j++
			}
			i++
		}
		return j == len(p)
	}

	l, r := 0, len(removable)+1
	for l < r {
		m := (r-l)>>1 + l
		if !isSubStr(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}
