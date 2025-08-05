package ac1

// 暴力模拟
func myAtoi(s string) int {
	res := 0
	i := 0
	sign := 1
	for ; i < len(s) && s[i] == ' '; i++ {
	}
	for j := i; j < len(s); j++ {
		if j == i {
			if s[j] == '-' {
				sign = -1
				continue
			}
			if s[j] == '+' {
				continue
			}
		}
		if s[j] < '0' || s[j] > '9' {
			break
		}
		res = res*10 + int(s[j]-'0')
		if sign == 1 && res > 1<<31-1 {
			res = 1<<31 - 1
			break
		} else if sign == -1 && res > 1<<31 {
			res = 1 << 31
			break
		}
	}
	return res * sign
}
