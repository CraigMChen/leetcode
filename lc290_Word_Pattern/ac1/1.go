package ac1

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	c2w := make(map[byte]string)
	w2c := make(map[string]byte)
	for i := 0; i < len(words); i++ {
		w, ok := c2w[pattern[i]]
		if ok && w != words[i] {
			return false
		}
		if !ok {
			c2w[pattern[i]] = words[i]
		}
		c, ok := w2c[words[i]]
		if ok && c != pattern[i] {
			return false
		}
		if !ok {
			w2c[words[i]] = pattern[i]
		}
	}
	return true
}
