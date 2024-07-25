package ac2

// 括号展开
// 用 sign 表示当前的数字在所有括号展开后的正负
// 用 stack 维护当前括号中的数字在所有括号展开后的正负
func calculate(s string) int {
	stack := []int{1}
	sign := 1
	ans := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}
