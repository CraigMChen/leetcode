package ac3

// KMP 算法，空间优化
// s := needle + "#" + heystack
// 对于任意的 i ∈ [0, len(s))，都有 pi[i] <= len(needle)
// 所以只需要存储 needle 的前缀数组和前一个前缀数组即可
func strStr(haystack string, needle string) int {
	pi := make([]int, len(needle))
	for i := 1; i < len(needle); i++ {
		j := pi[i-1]
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			pi[i] = j + 1
		} else {
			pi[i] = 0
		}
	}

	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j = j + 1
		} else {
			j = 0
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}
