package ac5

// 二分查找，左闭右闭区间
func longestCommonPrefix(strs []string) string {
	isLcp := func(s string) bool {
		for i := 0; i < len(strs); i++ {
			for j := 0; j < len(s); j++ {
				if j >= len(strs[i]) || strs[i][j] != s[j] {
					return false
				}
			}
		}
		return true
	}

	l, r := 0, len(strs[0])
	for l <= r {
		m := (r-l)>>1 + l
		if isLcp(strs[0][:m]) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return strs[0][:l-1]
}
