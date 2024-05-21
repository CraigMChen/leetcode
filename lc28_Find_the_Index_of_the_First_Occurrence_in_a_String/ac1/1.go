package ac1

// 暴力
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		j, k := 0, i
		for k < len(haystack) && j < len(needle) && haystack[k] == needle[j] {
			j++
			k++
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
