package ac2

// 朴素的 KMP 算法
func strStr(haystack string, needle string) int {
	s := needle + "#" + haystack
	pi := make([]int, len(s))
	for i := 1; i < len(pi); i++ {
		j := pi[i-1]
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			pi[i] = j + 1
		} else {
			pi[i] = 0
		}
		if pi[i] == len(needle) {
			return i - 2*len(needle)
		}
	}
	return -1
}
