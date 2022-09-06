package ac1

// 栈
// 用辅助栈来保存括号从而判断当前括号是否需要删除
// 用数组来保存需要删除的字符的下标
// 从头遍历 s，将不需要删除的字符构造成一个新串
func minRemoveToMakeValid(s string) string {
	var (
		stack  []byte
		index  []int
		res    []byte
		remove = make(map[int]struct{})
	)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, '(')
			index = append(index, i)
		} else if s[i] == ')' {
			n := len(stack)
			if n > 0 && stack[n-1] == '(' {
				stack = stack[:n-1]
				index = index[:n-1]
			} else {
				remove[i] = struct{}{}
			}
		}
	}
	for i := 0; i < len(index); i++ {
		remove[index[i]] = struct{}{}
	}
	for i := 0; i < len(s); i++ {
		if _, ok := remove[i]; !ok {
			res = append(res, s[i])
		}
	}
	return string(res)
}
