package ac3

// 分治
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return lcp(strs, 0, len(strs) - 1)
}

func lcp(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	} else if r == l + 1 {
		return getLCP(strs[l], strs[r])
	}
	m := (l + r) / 2
	return getLCP(lcp(strs, l, m), lcp(strs, m + 1, r))
}

func getLCP(a, b string) string {
	i := 0
	for ; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			break
		}
	}
	return a[:i]
}
