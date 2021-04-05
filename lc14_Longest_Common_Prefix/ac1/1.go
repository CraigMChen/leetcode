package ac1

// 横向扫描
func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	ans := strs[0]
	for i := 1; i < l; i++ {
		l1 := len(ans)
		l2 := len(strs[i])
		j := 0
		for ; j < l1 && j < l2; j++ {
			if ans[j] != strs[i][j]	{
				ans = ans[:j]
				break
			}
		}
		ans = ans[:j]
	}
	return ans
}
