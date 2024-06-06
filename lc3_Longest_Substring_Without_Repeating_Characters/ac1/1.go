package ac1

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	set := make([]bool, 128)
	l, r := 0, 0 // 左闭右开区间
	ans := 0
	for r < len(s) {
		if set[s[r]] {
			set[s[l]] = false
			l++
		} else {
			set[s[r]] = true
			r++
		}
		ans = max(ans, r-l)
	}
	return ans
}
