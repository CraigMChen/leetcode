package ac1

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s) && j >= 0; {
		a, b := s[i], s[j]
		if a >= 'A' && a <= 'Z' {
			a += 32
		} else if !(a >= 'a' && a <= 'z' || a >= '0' && a <= '9') {
			i++
			continue
		}
		if b >= 'A' && b <= 'Z' {
			b += 32
		} else if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9') {
			j--
			continue
		}
		if a != b {
			return false
		}
		i++
		j--
	}
	return true
}
