package ac1

// 用哈希表保存每个长度为 10 的子串的出现次数
func findRepeatedDnaSequences(s string) []string {
	counter := make(map[string]int)
	var res []string
	for i := 9; i < len(s); i++ {
		subStr := s[i-9 : i+1]
		counter[subStr]++
		if counter[subStr] == 2 {
			res = append(res, subStr)
		}
	}
	return res
}
