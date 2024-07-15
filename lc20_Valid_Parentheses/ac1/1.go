package ac1

// 栈
func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		n := len(stack)
		if n == 0 ||
			s[i] == ')' && stack[n-1] != '(' ||
			s[i] == ']' && stack[n-1] != '[' ||
			s[i] == '}' && stack[n-1] != '{' {
			return false
		}
		stack = stack[:n-1]
	}
	return len(stack) == 0
}
