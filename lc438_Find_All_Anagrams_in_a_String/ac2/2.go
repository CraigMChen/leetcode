package ac2

// 优化的滑动窗口
// 维护一个 count 数组，用于表示窗口中的子串和 p 中每个字母出现次数的差
// 再维护一个 diff，表示窗口中的子串和 p 中出现次数不一致的字母的个数，即 count[i] 不为 0 的个数
// 窗口滑动的过程中不断维护 diff，当 diff 位 0 时窗口子串即为 p 的异位词
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	count := [26]int{}
	for i := 0; i < len(p); i++ {
		count[s[i]-'a']++
		count[p[i]-'a']--
	}

	diff := 0
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			diff++
		}
	}

	var ans []int
	for i := 0; i < len(s)-len(p); i++ {
		if diff == 0 {
			ans = append(ans, i)
		}
		// 窗口左边界向右移动
		ch := s[i] - 'a'
		if count[ch] == 1 {
			diff--
		} else if count[ch] == 0 {
			diff++
		}
		count[ch]--

		// 窗口右边界向右移动
		ch = s[i+len(p)] - 'a'
		if count[ch] == -1 {
			diff--
		} else if count[ch] == 0 {
			diff++
		}
		count[ch]++
	}
	if diff == 0 {
		ans = append(ans, len(s)-len(p))
	}
	return ans
}
