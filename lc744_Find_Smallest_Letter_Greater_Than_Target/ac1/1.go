package ac1

// 二分
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)
	for l < r {
		m := (r - l) >> 1 + l
		if letters[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}
	if l == len(letters) {
		l = 0
	}
	return letters[l]
}
