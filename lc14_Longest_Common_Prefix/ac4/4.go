package ac4

// 二分查找，左闭右开区间
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := 200
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	l, r := 0, minLen + 1
	for l < r {
		m := (r-l)/2 + l
		if isLCP(m, strs) {
			l = m + 1
		} else {
			r = m
		}
	}
	if l == 0 {
		return ""
	}
	return strs[0][:l-1]
}

func isLCP(l int, strs []string) bool {
	for i := 0; i < l; i++ {
		for j := 0; j < len(strs); j++ {
			if j != len(strs)-1 && strs[j][i] != strs[j+1][i] {
				return false
			}
		}
	}
	return true
}
