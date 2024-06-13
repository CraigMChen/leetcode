package ac2

// 滑动窗口
// 思路与 lc438.ac2 类似
// wordLen 为 words 中单词的长度
// 区别是本题需要将 s 分成若干个长度为 wordLen 的单词，有 wordLen 种分法
// 分别为把 s 前 i (i ∈ [0, wordLen))个元素删去后，剩余的字母按照顺序组成若干个长度为 wordLen 的单词，多余的也删去
// 对于每一种分法都做一次滑动窗口
func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	totLen := wordLen * len(words)
	var ans []int
	for i := 0; i < wordLen && i+totLen <= len(s); i++ {
		count := make(map[string]int)
		for j := 0; j < len(words); j++ {
			count[s[i+j*wordLen:i+(j+1)*wordLen]]++
			count[words[j]]--
		}
		diff := 0
		for _, val := range count {
			if val != 0 {
				diff++
			}
		}

		if diff == 0 {
			ans = append(ans, i)
		}
		for j := i; j+totLen+wordLen <= len(s); j += wordLen {
			n := count[s[j:j+wordLen]]
			if n == 1 {
				diff--
			} else if n == 0 {
				diff++
			}
			count[s[j:j+wordLen]]--

			n = count[s[j+totLen:j+totLen+wordLen]]
			if n == -1 {
				diff--
			} else if n == 0 {
				diff++
			}
			count[s[j+totLen:j+totLen+wordLen]]++

			if diff == 0 {
				ans = append(ans, j+wordLen)
			}
		}
	}
	return ans
}
