package ac2

// 计数
// 将字符串中每个字母出现次数作为哈希表的键
func groupAnagrams(strs []string) [][]string {
	hash := func(str string) [26]int8 {
		var res [26]int8
		for i := 0; i < len(str); i++ {
			res[str[i]-'a']++
		}
		return res
	}
	var (
		res [][]string
		mp  = make(map[[26]int8]int)
	)
	for i := 0; i < len(strs); i++ {
		h := hash(strs[i])
		id, ok := mp[h]
		if !ok {
			res = append(res, make([]string, 0))
			id = len(res) - 1
			mp[h] = id
		}
		res[id] = append(res[id], strs[i])
	}
	return res
}
