package ac2

// 纵向扫描
// 时间复杂度 O(mn)
// 空间复杂度 O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; ; i++ {
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
}
