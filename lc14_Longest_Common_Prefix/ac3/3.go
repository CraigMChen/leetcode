package ac3

// 分治
func longestCommonPrefix(strs []string) string {
	lcp := func(a, b string) string {
		i := 0
		for ; i < len(a) && i < len(b) && a[i] == b[i]; i++ {
		}
		return a[:i]
	}

	var nLcp func(l, r int) string
	nLcp = func(l, r int) string {
		if l == r {
			return strs[l]
		} else if l+1 == r {
			return lcp(strs[l], strs[r])
		}
		m := (r-l)>>1 + l
		return lcp(nLcp(l, m), nLcp(m, r))
	}
	return nLcp(0, len(strs)-1)
}
