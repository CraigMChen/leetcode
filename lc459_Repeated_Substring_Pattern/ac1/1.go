package ac1

// 暴力模拟
func repeatedSubstringPattern(s string) bool {
	for i := 0; i*2 <= len(s); i++ {
		if i == len(s)-1 || len(s)%(i+1) != 0 {
			continue
		}
		j := i + 1
		for ; j < len(s) && s[j] == s[j%(i+1)]; j++ {
		}
		if j == len(s) {
			return true
		}
	}
	return false
}
