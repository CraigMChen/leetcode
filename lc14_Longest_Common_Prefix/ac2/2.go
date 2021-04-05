package ac2

// 纵向扫描
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; ; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if i >= len(strs[j]) || i >= len(strs[j+1]) {
				return strs[0][:i]
			}
			if strs[j][i] != strs[j+1][i] {
				return strs[0][:i]
			}
		}
	}
}
