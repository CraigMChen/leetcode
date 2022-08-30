package ac2

const L = 10

var mp = map[byte]int{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

// 位运算
// 一个字母可以用 2 位二进制数表示，长度为 10 的字符串可以表示为 20 位的二进制数，用一个 int 保存足够
// 用这个 int 值作为哈希表的键
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}
	var (
		counter = make(map[int]int)
		res     []string
		key     = 0
	)
	for i := 0; i < L-1; i++ {
		key = key<<2 | mp[s[i]]
	}
	for i := 0; i <= len(s)-L; i++ {
		key = (key<<2 | mp[s[i+L-1]]) & (1<<(L*2) - 1)
		counter[key]++
		if counter[key] == 2 {
			res = append(res, s[i:i+L])
		}
	}
	return res
}
