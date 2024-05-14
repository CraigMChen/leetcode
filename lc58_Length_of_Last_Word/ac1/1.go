package ac1

// 正向遍历
func lengthOfLastWord(s string) int {
	pre, cur := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && cur != 0 {
			pre = cur
			cur = 0
		} else if s[i] != ' ' {
			cur++
		}
	}
	if cur != 0 {
		return cur
	}
	return pre
}
