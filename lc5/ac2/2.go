package ac2

// 从中心往两边扩展
func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	palindromeLen := 1
	start, end := 0, 0
	for i := 0; i < l; i++ {
		for j := 0; i-j >= 0 && i+j < l; j++ {
			if s[i-j] != s[i + j] {
				break
			}
			if 2 * j + 1 > palindromeLen {
				palindromeLen = 2 * j + 1
				start = i - j
				end = i + j
			}
		}
		if i != l - 1 && s[i] == s[i + 1] {
			for j := 0; i - j >=0 && i + j + 1 < l; j++ {
				if s[i - j] != s[i + j + 1] {
					break
				}
				if 2 * j + 2 > palindromeLen {
					palindromeLen = 2 * j + 2
					start = i - j
					end = i + j + 1
				}
			}
		}
	}
	return s[start:end+1]
}
