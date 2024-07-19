package ac1

import "strconv"

// 栈
func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		if len(token) == 1 && (token[0] == '+' ||
			token[0] == '-' || token[0] == '*' ||
			token[0] == '/') {
			var a, b int
			b = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
			continue
		}
		val, _ := strconv.Atoi(token)
		stack = append(stack, val)
	}
	return stack[0]
}
