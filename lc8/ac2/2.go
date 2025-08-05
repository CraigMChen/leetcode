package ac2

// 有限状态机
const (
	START = iota
	SIGN
	NUMBER
	END
)

var automaton = [3][4]int{
	// space, operator, number, other
	{START, SIGN, NUMBER, END}, // START
	{END, END, NUMBER, END},    // SIGN
	{END, END, NUMBER, END},    // NUMBER
}

func myAtoi(s string) int {
	res := 0
	curState := START
	sign := 1
	for i := 0; i < len(s); i++ {
		if curState == END {
			break
		}
		id := 0
		switch {
		case s[i] == ' ':
			id = 0
		case s[i] == '-' || s[i] == '+':
			id = 1
		case s[i] >= '0' && s[i] <= '9':
			id = 2
		default:
			id = 3
		}
		curState = automaton[curState][id]
		if curState == SIGN && s[i] == '-' {
			sign = -1
		}
		if curState == NUMBER && s[i] >= '0' && s[i] <= '9' {
			res = res*10 + int(s[i]-'0')
			if sign == 1 && res > 1<<31-1 {
				return 1<<31 - 1
			} else if sign == -1 && res > 1<<31 {
				return -1 << 31
			}
		}
	}
	return res * sign
}
