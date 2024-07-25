package ac1

import "strconv"

// 栈 模拟
func calculate(s string) int {
	var stack []interface{}
	top := func() interface{} {
		return stack[len(stack)-1]
	}
	pop := func() interface{} {
		o := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return o
	}
	// 统一处理加减法
	cal := func(s1 int) {
		if len(stack) == 0 {
			stack = append(stack, s1)
		} else {
			op := top().(byte)
			switch op {
			case '+':
				pop()
				s2 := pop().(int)
				stack = append(stack, s1+s2)
			case '-':
				pop()
				if len(stack) == 0 {
					stack = append(stack, -s1)
				} else {
					s2, ok := top().(int)
					if ok {
						pop()
						stack = append(stack, s2-s1)
					} else {
						stack = append(stack, -s1)
					}
				}
			case '(':
				stack = append(stack, s1)
			}
		}
	}

	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}
		if s[i] == '(' || s[i] == '+' || s[i] == '-' {
			stack = append(stack, s[i])
			i++
		} else if s[i] == ')' {
			s1 := pop().(int)
			pop()
			cal(s1)
			i++
		} else {
			j := i + 1
			for ; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
			}
			s1, _ := strconv.Atoi(s[i:j])
			cal(s1)
			i = j
		}
	}
	return stack[len(stack)-1].(int)
}
