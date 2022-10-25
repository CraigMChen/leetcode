package ac2

// 从中心往两边扩展
func longestPalindrome(s string) string {
	expand := func(start, end int) (int, int) {
		for ; start >= 0 && end < len(s) && s[start] == s[end]; start, end = start-1, end+1 {
		}
		return start + 1, end - 1
	}
	maxLen := 0
	res := ""
	for i := 0; i < len(s); i++ {
		start1, end1 := expand(i, i)
		start2, end2 := expand(i, i+1)
		if end1-start1+1 > maxLen {
			maxLen = end1 - start1 + 1
			res = s[start1 : end1+1]
		}
		if end2-start2+1 > maxLen {
			maxLen = end2 - start2 + 1
			res = s[start2 : end2+1]
		}
	}
	return res
}
