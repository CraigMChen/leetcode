package ac2

// 反向遍历
func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if ans == 0 {
				continue
			} else {
				return ans
			}
		} else {
			ans++
		}
	}
	return ans
}
