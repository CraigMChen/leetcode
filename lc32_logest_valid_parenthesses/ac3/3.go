package ac3

// 计数
// 用 left 和 right 分别表示遍历过程中遇到的左括号和右括号的数量
// 先从左往右遍历
// 当 right > left 时清理计数值
// 当 right == left 时，表明以当前下标结尾的长度为 left + right 的串为合法括号串，更新答案
// 然后从右往左遍历
// 当 left > right 是清理计数值
// 当 right == left 时，更新答案
// 时间复杂度 O(n)
// 空间复杂的 O(1)
func longestValidParentheses(s string) int {
	left, right := 0, 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if right > left {
			left, right = 0, 0
		} else if right == left {
			ans = max(ans, left+right)
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if left > right {
			left, right = 0, 0
		} else if right == left {
			ans = max(ans, left+right)
		}
	}
	return ans
}
