package ac1

// 滑动窗口
// 与 lc30/lc438 的 ac2 思想类似
// 用 map count 统计 t 中字母的出现次数与窗口中字母出现次数的差（只需考虑 t 中出现过的字母）
// 区别是，用 diff 表示 t 中的出现次数大于窗口中出现次数的字母的个数
// 即 count 中大于 0 的元素数量
// 当 diff > 0 时，说明当前窗口表示的子串不符合要求，需要扩大，即将右边界向右移动
// 当 diff <= 0 时，说明当前窗口表示的子串符合要求，将左边界向右移动，并更新答案
func minWindow(s string, t string) string {
	count := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		count[t[i]]++
	}

	diff := 0
	for _, c := range count {
		if c > 0 {
			diff++
		}
	}

	ans := ""
	minLen := 100001
	l, r := 0, 0
	for r <= len(s) {
		if diff > 0 { // diff 大于 0，说明此时窗口表示的子串不符合要求，需要将右边界右移
			if r == len(s) { // r 已经到最右了，但是窗口还是不符合要求
				break
			}
			if _, ok := count[s[r]]; ok {
				count[s[r]]--
				if count[s[r]] == 0 { // count[s[r]] 从 1 变成 0，t 中出现次数大于窗口中出现次数的字母个数变少了
					diff--
				}
			}
			r++
		} else { // diff 小于等于 0，说明此时窗口表示的子串符合要求，尝试将左边界右移来使子串尽可能短
			if r-l < minLen {
				ans = s[l:r]
				minLen = r - l
			}
			if _, ok := count[s[l]]; ok {
				count[s[l]]++
				if count[s[l]] == 1 { // count[s[l]] 从 0 变成 1，t 中出现次数大于窗口个中出现次数的字母个数变多了
					diff++
				}
			}
			l++
		}
	}
	return ans
}
