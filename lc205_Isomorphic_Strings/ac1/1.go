package ac1

func isIsomorphic(s string, t string) bool {
	s2t := make([]byte, 128)
	t2s := make([]byte, 128)
	for i := 0; i < len(s); i++ {
		if s2t[s[i]] != 0 && s2t[s[i]] != t[i] {
			return false
		} else {
			s2t[s[i]] = t[i]
		}
		if t2s[t[i]] != 0 && t2s[t[i]] != s[i] {
			return false
		} else {
			t2s[t[i]] = s[i]
		}
	}
	return true
}
