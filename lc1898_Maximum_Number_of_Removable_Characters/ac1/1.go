package ac1

const maxInt = int(^uint(0) >> 1)

// 二分查找
// 用动态规划的方法判断子序列，类似 lc392
func maximumRemovals(s string, p string, removable []int) int {
	removed := make([]int, len(s)+1)
	for i := 0; i < len(removed); i++ {
		removed[i] = maxInt
	}
	for i := 0; i < len(removable); i++ {
		removed[removable[i]+1] = i
	}

	loc := [26]int{}
	for i := 0; i < 26; i++ {
		loc[i] = -1
	}
	nextLoc := make([][26]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		nextLoc[i+1] = loc
		loc[s[i]-'a'] = i + 1
	}
	nextLoc[0] = loc

	isSubStr := func(k int) bool {
		i, j := 0, 0
		for j < len(p) {
			i = nextLoc[i][p[j]-'a']
			if i == -1 {
				break
			}
			if removed[i] < k {
				continue
			}
			j++
		}
		return i != -1
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
