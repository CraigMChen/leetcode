package ac1

// 暴力模拟
func myAtoi(s string) int {
	var ans int64 = 0
	l := len(s)
	isPositive := 0
	getNum := false
	for i := 0; i < l; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			ans = ans*10 + int64(s[i]-'0')
			getNum = true
			if ans < (-1<<31)*10 || ans > (1<<31-1)*10 {
				break
			}
		} else {
			if getNum {
				break
			}
			if s[i] == '-' && isPositive == 0 {
				getNum = true
				isPositive = -1
			} else if s[i] == '+' && isPositive == 0 {
				getNum = true
				isPositive = 1
			} else if s[i] == ' ' {
				continue
			} else {
				break
			}
		}
	}
	if isPositive != 0 {
		ans *= int64(isPositive)
	}
	if ans < -1<<31 {
		ans = -1<<31
	} else if ans > 1<<31-1 {
		ans = 1<<31-1
	}
	return int(ans)
}
