package ac1

// 滑动窗口
// 分别统计窗口内和p中每个字母的出现次数
// 当两个数组相同时窗口表示的子串符合要求
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	countP := [26]int{}
	countS := [26]int{}
	for i := 0; i < len(p); i++ {
		countP[p[i]-'a']++
		countS[s[i]-'a']++
	}

	var ans []int
	for i := 0; i < len(s)-len(p); i++ {
		if countS == countP {
			ans = append(ans, i)
		}
		countS[s[i]-'a']--
		countS[s[i+len(p)]-'a']++
	}
	if countS == countP {
		ans = append(ans, len(s)-len(p))
	}
	return ans
}
