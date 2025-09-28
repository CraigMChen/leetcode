package ac3

// 分治
// 时间复杂度 O(mn)
// 空间复杂度 O(mlogn)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	m := len(strs) >> 1
	l := longestCommonPrefix(strs[:m])
	r := longestCommonPrefix(strs[m:])
	return lcp(l, r)
}

func lcp(str1, str2 string) string {
	i := 0
	for ; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			break
		}
	}
	return str1[:i]
}
