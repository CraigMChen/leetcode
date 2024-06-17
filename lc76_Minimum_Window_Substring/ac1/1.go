package ac1

// 滑动窗口
// 与 lc30/lc438 的 ac2 思想类似
// 用 map count 统计当前子串和 t 的每个字母的数量差（只需考虑 t 中出现过的字母）
// 用 diff 表示当前子串中出现次数大于等于 t 中出现次数的字母个数
// 当 diff > 0 时，说明当前子串不符合要求，需要扩大，即将右边界向右移动
// 当 diff <= 0 时，说明当前子串符合要求，将左边界向右移动
func minWindow(s string, t string) string {
	count := make(map[byte]int)
	diff := 0
	for i := 0; i < len(t); i++ {
		_, ok := count[t[i]]
		if !ok {
			diff++
		}
		count[t[i]]++
	}

	l, r := 0, 0 // 表示结果的左闭右开区间
	for i, j := 0, 0; j <= len(s); {
		if diff > 0 { // 此时，当前子串中的字母不够，需要将右边界右移
			if j == len(s) {
				break
			}
			_, ok := count[s[j]]
			if ok { // 只需考虑 t 中出现过的字母
				if count[s[j]] == 1 {
					diff--
				}
				count[s[j]]--
			}
			j++
		} else { // 此时，当前子串中的字母足够，需要将左边界右移
			if r-l == 0 || j-i <= r-l {
				l, r = i, j
			}
			_, ok := count[s[i]]
			if ok {
				if count[s[i]] == 0 {
					diff++
				}
				count[s[i]]++
			}
			i++
		}
	}
	return s[l:r]
}
