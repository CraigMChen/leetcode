package ac2

// 有限状态机
func myAtoi(s string) int {
	table := map[string][]string{
		"start":  {"start", "signed", "number", "end"},
		"signed": {"end", "end", "number", "end"},
		"number": {"end", "end", "number", "end"},
		"end":    {"end", "end", "end", "end"},
	}

	var ans int64 = 0
	sign := 1
	status := "start"
	l := len(s)
	for i := 0; i < l; i++ {
		status = table[status][getCol(s[i])]
		if status == "signed" && s[i] == '-' {
			sign = -1
		} else if status == "number" {
			ans = ans*10 + int64(s[i]-'0')
			if sign == 1 && ans > 1<<31-1 {
				ans = 1<<31 - 1
				break
			} else if sign == -1 && ans > 1<<31 {
				ans = 1 << 31
				break
			}
		} else if status == "end" {
			break
		}
	}
	return int(ans) * sign
}

func getCol(c byte) int {
	if c == ' ' {
		return 0
	} else if c == '+' || c == '-' {
		return 1
	} else if c >= '0' && c <= '9' {
		return 2
	}
	return 3
}
