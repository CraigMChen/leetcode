package ac2

// KMP 算法
// 先求出前缀数组 pi
// 若 len(s) 是 len(s) - pi[len(s)-1] 的倍数，则 s 一定满足条件
func repeatedSubstringPattern(s string) bool {
	l := len(s)
	pi := make([]int, l)
	for i := 1; i < l; i++ {
		j := pi[i-1]
		for j != 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			pi[i] = j + 1
		}
	}
	return pi[l-1] != 0 && l%(l-pi[l-1]) == 0
}
