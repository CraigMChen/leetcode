package ac2

// 栈
// 用栈底元素表示已遍历元素中最后一个没有匹配的右括号的下标
// 其他元素表示未匹配的左括号的下标
// 初始时栈中放入一个 -1，表述虚拟的右括号
// 遍历 s，若 s[i] 为左括号则将 i 入栈
// 否则，将栈顶元素出栈
// 栈顶元素出栈后若栈的长度为空，表明刚才出栈的元素为右括号，不能与当前右括号匹配
// 则将当前右括号的下标入栈，表示最后一个没有匹配的右括号的下标
// 栈顶元素出栈后若栈的长度不为空，表明刚才出栈的元素为左括号，与当前右括号匹配
// 则当前下标减去栈顶元素即为以当前下标结尾的有效括号的长度
// 遍历过程中得到的最大值即为最终答案
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func longestValidParentheses(s string) int {
	stack := make([]int, 0, len(s))
	stack = append(stack, -1)
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			ans = max(ans, i-stack[len(stack)-1])
		} else {
			stack = append(stack, i)
		}
	}
	return ans
}
