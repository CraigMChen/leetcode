package ac1

import "strings"

// 回溯
func restoreIpAddresses(s string) []string {
	var ans []string
	var f func([]string, int)
	f = func(cur []string, i int) {
		if len(cur) == 4 && i < len(s) {
			return
		}
		if i >= len(s) {
			if len(cur) == 4 {
				ans = append(ans, strings.Join(cur, "."))
			}
			return
		}

		for j := i; j < i+3 && j < len(s); j++ {
			seg := s[i : j+1]
			if seg[0] == '0' && len(seg) > 1 {
				break
			}
			if len(seg) == 3 && seg > "255" {
				break
			}
			cur = append(cur, seg)
			f(cur, j+1)
			cur = cur[:len(cur)-1]
		}
	}
	f([]string{}, 0)
	return ans
}
