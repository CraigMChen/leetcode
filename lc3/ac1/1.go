package ac1

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	bitSet := make([]bool, 256)
	ans, head, tail := 0, 0, 0
	for {
		if bitSet[s[head]] {
			bitSet[s[tail]] = false
			tail++
		} else {
			bitSet[s[head]] = true
			head++
		}
		if head - tail > ans {
			ans = head - tail
		}
		if head >= l {
			break
		}
	}
	return ans
}
