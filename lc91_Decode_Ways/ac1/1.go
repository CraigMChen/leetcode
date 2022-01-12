package ac1

// 动态规划
// 设 dp[i] 为 s[0:i+1] 能被解码的方法数
// 对于 s[i]，分两种情况讨论：
// 若 s[i] 能被解码，即 s[i] != 0，则这种情况下的方法数为 dp[i-1]，否则为 0
// 若 s[i-1] 和 s[i] 组成的字符串能被解码，即 s[i-1] != 0 且 1 <= s[i-1] * 10 + s[i] <= 26，则这种情况下的方法数为 dp[i-1]，否则为 0
// dp[i] 为上面两种情况的方法数之和
// 边界条件： dp[-1] = 1 （空串可以被解码为空串），若 s[0] 可被解码，则 dp[0] = 1，否则 dp[0] = 0
// dp[len(s) - 1]即为答案
func numDecodings(s string) int {
	x, y := 1, 0
	if canBeDecoded(s[0]) {
		y = 1
	}
	for i := 1; i < len(s); i++ {
		tmp := 0
		if canBeDecoded(s[i]) {
			tmp += y
		}
		if canBeDecoded(s[i-1], s[i]) {
			tmp += x
		}
		x, y = y, tmp
	}
	return y
}

func canBeDecoded(code ...byte) bool {
	if code[0] == '0' {
		return false
	}
	sum := byte(0)
	for i := 0; i < len(code); i++ {
		sum *= 10
		sum += code[i] - '0'
	}
	if int8(sum) >= 1 && int8(sum) <= 26 {
		return true
	}
	return false
}
