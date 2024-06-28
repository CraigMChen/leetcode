package ac1

func isAnagram(s string, t string) bool {
	counter := [26]int{}
	for i := 0; i < len(s) || i < len(t); i++ {
		if i < len(s) {
			counter[s[i]-'a']++
		}
		if i < len(t) {
			counter[t[i]-'a']--
		}
	}
	for i := 0; i < len(counter); i++ {
		if counter[i] != 0 {
			return false
		}
	}
	return true
}
