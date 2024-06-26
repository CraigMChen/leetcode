package ac1

// 横向扫描
func longestCommonPrefix(strs []string) string {
	l := len(strs[0])
	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < l && j < len(strs[i]) && strs[i][j] == strs[0][j]; j++ {
		}
		l = j
	}
	return strs[0][:l]
}
