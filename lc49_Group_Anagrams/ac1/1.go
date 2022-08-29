package ac1

import "sort"

// 排序
// 将字符串排序后作为哈希表的键
func groupAnagrams(strs []string) [][]string {
	hash := func(str string) string {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		return string(s)
	}
	var (
		res [][]string
		mp  = make(map[string]int)
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
