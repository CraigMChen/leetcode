package ac1

// 暴力模拟
// wordLen 表示 words 中单词的长度
// totLen 表示 words 中所有单词的长度之和
// count 表示 words 中每个单词的出现次数
// 遍历 s，对于每个 s 的长度为 totLen 的子串，用 countS 统计该子串按顺序分为长度为 wordLen 的每个单词的出现次数
// 比较 count 和 countS 两个 map 中元素是否相同，相同则当前子串符合要求
func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	totLen := wordLen * len(words)
	count := make(map[string]int)
	for i := 0; i < len(words); i++ {
		count[words[i]]++
	}

	var ans []int
	for i := 0; i+totLen <= len(s); i++ {
		countS := make(map[string]int)
		for j := 0; j < len(words); j++ {
			countS[s[i+j*wordLen:i+(j+1)*wordLen]]++
		}
		diff := false
		for k, v := range countS {
			if c := count[k]; c != v {
				diff = true
			}
		}
		for k, v := range count {
			if c := countS[k]; c != v {
				diff = true
			}
		}
		if !diff {
			ans = append(ans, i)
		}
	}
	return ans
}
